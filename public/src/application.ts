(function () {
    type TypeaheadResult = Promise<Array<string>>;

    function wireTypeahead(
        inputId: string,
        listId: string,
        dataCallback: (value: string) => TypeaheadResult
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
        dataCallback: (value: string) => TypeaheadResult
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

    async function getList(path: string, responseKey: string, body: object): TypeaheadResult {
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

    async function vendorCallback(vendor: string): TypeaheadResult {
        const body = {
            vendor: vendor
        };

        return getList("/vendor", "vendors", body)
    }

    async function productCallback(product: string): TypeaheadResult {
        const vendor = document.getElementById("vendor") as HTMLInputElement;
        if (vendor === null) {
            return [];
        }

        const body = {
            vendor: vendor.value,
            name: product
        };

        return getList("/product", "products", body);
    }

    async function versionCallback(version: string): TypeaheadResult {
        const vendor = document.getElementById("vendor") as HTMLInputElement;
        if (vendor === null) {
            return [];
        }

        const product = document.getElementById("product") as HTMLInputElement;
        if (product === null) {
            return [];
        }

        const body = {
            vendor: vendor.value,
            name: product.value,
            version: version
        };

        return getList("/version", "versions", body);
    }

    wireTypeahead("vendor", "vendor-list", vendorCallback);
    wireTypeahead("product", "product-list", productCallback);
    wireTypeahead("version", "version-list", versionCallback);
})();
