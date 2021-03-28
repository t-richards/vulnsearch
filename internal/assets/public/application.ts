type TypeaheadResult = Promise<string[]>;
type InputId = "vendor" | "name" | "version";

// Returns the value of the input or an empty string.
const getInputValue = (id: InputId): string => {
    const el = document.getElementById(id) as HTMLInputElement | null;
    if (el === null) { return ""; }
    return el.value;
}

// Creates an HTML element with additional properties
const h = <K extends keyof HTMLElementTagNameMap>(tagName: K, props: Record<string, unknown>): HTMLElementTagNameMap[K] => {
    return Object.assign(document.createElement(tagName), props);
}

/**
 * Wires up a typeahead "component" using a native HTML input and datalist element.
 *
 * @param inputId The id of the input element to listen for input events
 * @param listId The associated datalist element to be updated
 * @param dataCallback A callback function which is expected to return the set of datalist options
 */
const wireTypeahead = (
    inputId: InputId,
    listId: string,
    dataCallback: () => TypeaheadResult
): void => {
    const input = document.getElementById(inputId);
    input?.addEventListener("change", createInputHandler(listId, dataCallback));
    input?.addEventListener("input", createInputHandler(listId, dataCallback));
}

const createInputHandler = (
    listId: string,
    dataCallback: () => TypeaheadResult
) => {
    return async function (_evt: Event): Promise<void> {
        const data = await dataCallback();
        const newList = h("datalist", { id: listId });

        for (const item of data) {
            newList.appendChild(
                h("option", { value: item, innerHTML: item })
            );
        }

        const originalList = document.getElementById(listId);
        originalList?.parentNode?.replaceChild(newList, originalList);
    };
}

const getResults = async (path: string, responseKey: string, body: Record<string, unknown>): TypeaheadResult => {
    try {
        const response = await fetch(path, {
            method: "POST",
            body: JSON.stringify(body),
            headers: {
                "Content-Type": "application/json"
            }
        });
        const jsonResponse = await response.json();
        const data = jsonResponse[responseKey] as string[];
        return data;
    } catch (err) {
        console.error(err);
        return [];
    }
}

const vendorCallback = async (): TypeaheadResult => {
    return getResults(
        "/vendor",
        "vendors",
        { vendor: getInputValue("vendor") }
    );
}

const nameCallback = async (): TypeaheadResult => {
    return getResults(
        "/product",
        "products",
        { vendor: getInputValue("vendor"), name: getInputValue("name") }
    );
}

const versionCallback = async (): TypeaheadResult => {
    return getResults(
        "/version",
        "versions",
        { vendor: getInputValue("vendor"), name: getInputValue("name"), version: getInputValue("version") }
    );
}

const init = () => {
    // This code only runs on / with the search form
    if (document.getElementById("search") === null) {
        return;
    }

    wireTypeahead("vendor", "vendor-list", vendorCallback);
    wireTypeahead("name", "name-list", nameCallback);
    wireTypeahead("version", "version-list", versionCallback);
}

init();
