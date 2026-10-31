// Delay between the last keystroke and fetching suggestions.
const debounceMs = 150

// Suggestions endpoint for each input, keyed by input name.
const endpoints: Record<string, string> = {
  vendor: "/api/vendors",
  name: "/api/products",
  version: "/api/versions",
}

const inputs = [...(document.getElementById("search")?.querySelectorAll("input") ?? [])]

// Every endpoint receives the whole form.
const query = () => new URLSearchParams(inputs.map((input) => [input.name, input.value]))

const sleep = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms))

// Fills the datalist with suggestions from the endpoint, after a debounce.
const suggest = async (endpoint: string, list: HTMLDataListElement, signal: AbortSignal) => {
  await sleep(debounceMs)
  signal.throwIfAborted()

  const response = await fetch(`${endpoint}?${query()}`, { signal })
  if (!response.ok) throw new Error(`${endpoint}: ${response.status}`)

  const items: string[] = await response.json()
  list.replaceChildren(...items.map((item) => new Option(item, item)))
}

for (const input of inputs) {
  const endpoint = endpoints[input.name]
  const list = input.list
  if (!endpoint || !list) continue

  // Aborting cancels both a pending debounce and an in-flight request.
  let controller = new AbortController()

  const refresh = () => {
    controller.abort()
    const { signal } = (controller = new AbortController())
    suggest(endpoint, list, signal).catch((err) => {
      if (!signal.aborted) console.error(err)
    })
  }

  input.addEventListener("input", refresh)
  input.addEventListener("focus", refresh)
}
