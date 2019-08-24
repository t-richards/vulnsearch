(function () {
    type TypeaheadData = Promise<Array<string>>;

    function wireTypeahead(
        inputId: string,
        listId: string,
        dataCallback: (value: string) => TypeaheadData
    ): void {
        const input = document.getElementById(inputId);
        if (input === null) {
            console.error(`Failed to locate ${inputId}.`);
            return;
        }

        const handler = createInputHandler(listId, dataCallback);
        input.addEventListener("input", handler);
    }

    function createInputHandler(
        listId: string,
        dataCallback: (value: string) => TypeaheadData
    ) {
        return async function (evt: Event): Promise<void> {
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

    async function vendorCallback(vendor: string): TypeaheadData {
        const body = {
            vendor: vendor
        };

        try {
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
        } catch (err) {
            console.error(err);
            return [];
        }
    }

    async function productCallback(product_name: string): TypeaheadData {
        const vendor = document.getElementById("vendor") as HTMLInputElement;
        if (vendor === null) {
            return [];
        }

        const body = {
            vendor: vendor.value,
            name: product_name
        }

        try {
            const response = await fetch("/product", {
                method: "POST",
                body: JSON.stringify(body),
                headers: {
                    "Content-Type": "application/json"
                }
            });
            const jsonResponse = await response.json();
            const data = jsonResponse.products as Array<string>;
            return data;
        } catch (err) {
            console.error(err);
            return [];
        }
    }

    wireTypeahead("vendor", "vendor-list", vendorCallback);
    wireTypeahead("product", "product-list", productCallback)
})();
