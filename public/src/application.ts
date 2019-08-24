(function() {
    type TypeaheadData = Promise<Array<string>>;

    function wireTypeahead(
        inputId: string,
        listId: string,
        dataCallback: (value: string) => TypeaheadData
    ) {
        const input = document.getElementById(inputId);
        if (input !== null) {
            const handler = createInputHandler(listId, dataCallback);
            input.addEventListener("input", handler);
        }
    }

    function createInputHandler(
        listId: string,
        dataCallback: (value: string) => TypeaheadData
    ) {
        return async function(evt: Event): Promise<void> {
            if (evt.target === null) {
                return;
            }
            const target = evt.target as HTMLInputElement;
            const data = await dataCallback(target.value);
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
        };
    }

    async function vendorCallback(value: string): TypeaheadData {
        const body = {
            vendor: value
        };

        const response = await fetch("/vendor", {
            method: "POST",
            body: JSON.stringify(body),
            headers: {
                "Content-Type": "application/json"
            }
        });
        const jsonResponse = await response.json();
        const data = jsonResponse.vendors as Array<string>;
        return data;
    }

    wireTypeahead("vendor", "vendor-list", vendorCallback);
})();
