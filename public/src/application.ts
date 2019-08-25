(function () {
    type TypeaheadResult = Promise<Array<string>>;
    type InputId = 'vendor' | 'name' | 'version';
    interface ApplicationState {
        vendor: string,
        name: string,
        version: string
    }

    const state: ApplicationState = {
        vendor: '',
        name: '',
        version: '',
    };

    /**
     * Wires up a typeahead "component" using a native HTML input and datalist element.
     *
     * @param inputId The id of the input element to listen for input events
     * @param listId The associated datalist element to be updated
     * @param dataCallback A callback function which is expected to return the set of datalist options
     */
    function wireTypeahead(
        inputId: InputId,
        listId: string,
        dataCallback: () => TypeaheadResult
    ): void {
        const input = document.getElementById(inputId);
        if (input === null) {
            console.error(`Failed to locate ${inputId}.`);
            return;
        }

        input.addEventListener("input", createInputHandler(inputId, listId, dataCallback));
    }

    function createInputHandler(
        inputId: InputId,
        listId: string,
        dataCallback: () => TypeaheadResult
    ) {
        return async function (evt: Event): Promise<void> {
            const markerStart = "replace-list-start";
            const markerEnd = "replace-list-end";

            if (evt.target === null) {
                console.error("Event has no target!");
                return;
            }
            const target = evt.target as HTMLInputElement;
            state[inputId] = target.value;
            const data = await dataCallback();

            performance.mark(markerStart);
            const newList = document.createElement("datalist");
            newList.id = listId;

            data.forEach(item => {
                let opt = document.createElement("option");
                opt.value = item;
                opt.innerHTML = item;
                newList.appendChild(opt);
            });

            const originalList = document.getElementById(listId);
            if (originalList === null || originalList.parentNode === null) {
                console.error("Failed to replace original list.");
                return;
            }
            originalList.parentNode.replaceChild(newList, originalList);
            performance.mark(markerEnd);

            // Log perf
            performance.measure("Typeahead list replaced", markerStart, markerEnd);
            console.log(performance.getEntriesByType("measure")[0]);
            performance.clearMarks();
            performance.clearMeasures();
        };
    }

    async function getResults(path: string, responseKey: string, body: object): TypeaheadResult {
        try {
            const response = await fetch(path, {
                method: "POST",
                body: JSON.stringify(body),
                headers: {
                    "Content-Type": "application/json"
                }
            });
            const jsonResponse = await response.json();
            const data = jsonResponse[responseKey] as Array<string>;
            return data;
        } catch (err) {
            console.error(err);
            return [];
        }
    }

    async function vendorCallback(): TypeaheadResult {
        return getResults(
            "/vendor",
            "vendors",
            { vendor: state.vendor }
        );
    }

    async function nameCallback(): TypeaheadResult {
        return getResults(
            "/product",
            "products",
            { vendor: state.vendor, name: state.name }
        );
    }

    async function versionCallback(): TypeaheadResult {
        return getResults(
            "/version",
            "versions",
            { vendor: state.vendor, name: state.name, version: state.version }
        );
    }

    wireTypeahead("vendor", "vendor-list", vendorCallback);
    wireTypeahead("name", "name-list", nameCallback);
    wireTypeahead("version", "version-list", versionCallback);
})();
