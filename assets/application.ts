type InputId = "vendor" | "name" | "version"

// Delay between the last keystroke and fetching suggestions.
const debounceMs = 150

// Returns the value of the input or an empty string.
const getInputValue = (id: InputId): string => {
  const el = document.getElementById(id) as HTMLInputElement | null
  return el?.value ?? ""
}

// Fetches a list of suggestions, aborting any previous in-flight request.
const fetchSuggestions = async (path: string, params: Record<string, string>, signal: AbortSignal): Promise<string[]> => {
  const response = await fetch(`${path}?${new URLSearchParams(params)}`, { signal })
  if (!response.ok) {
    throw new Error(`${path}: ${response.status}`)
  }
  return await response.json() as string[]
}

/**
 * Wires up a typeahead "component" using a native HTML input and datalist element.
 *
 * @param inputId The id of the input element to listen for input events
 * @param listId The associated datalist element to be updated
 * @param path The suggestions endpoint
 * @param params Returns the query parameters for the suggestions endpoint
 */
const wireTypeahead = (
  inputId: InputId,
  listId: string,
  path: string,
  params: () => Record<string, string>,
): void => {
  const input = document.getElementById(inputId)
  const list = document.getElementById(listId)
  if (input === null || list === null) {
    return
  }

  let timer: number | undefined
  let controller: AbortController | undefined

  const update = async (): Promise<void> => {
    controller?.abort()
    controller = new AbortController()

    try {
      const items = await fetchSuggestions(path, params(), controller.signal)
      list.replaceChildren(...items.map((value) => new Option(value, value)))
    } catch (err) {
      if (!(err instanceof DOMException && err.name === "AbortError")) {
        console.error(err)
      }
    }
  }

  const schedule = (): void => {
    window.clearTimeout(timer)
    timer = window.setTimeout(update, debounceMs)
  }

  input.addEventListener("input", schedule)
  input.addEventListener("focus", schedule)
}

const init = (): void => {
  wireTypeahead("vendor", "vendor-list", "/api/vendors", () => ({
    vendor: getInputValue("vendor"),
  }))
  wireTypeahead("name", "name-list", "/api/products", () => ({
    vendor: getInputValue("vendor"),
    name: getInputValue("name"),
  }))
  wireTypeahead("version", "version-list", "/api/versions", () => ({
    vendor: getInputValue("vendor"),
    name: getInputValue("name"),
    version: getInputValue("version"),
  }))
}

init()
