"use strict";
{
    const state = {
        vendor: getInputValue("vendor"),
        name: getInputValue("name"),
        version: getInputValue("version")
    };
    function getInputValue(id) {
        const el = document.getElementById(id);
        if (el === null) {
            return "";
        }
        return el.value;
    }
    function wireTypeahead(inputId, listId, dataCallback) {
        const input = document.getElementById(inputId);
        if (input === null) {
            console.error(`Failed to locate ${inputId}.`);
            return;
        }
        input.addEventListener("change", createInputHandler(inputId, listId, dataCallback));
        input.addEventListener("input", createInputHandler(inputId, listId, dataCallback));
    }
    function createInputHandler(inputId, listId, dataCallback) {
        return async function (evt) {
            const markerStart = "replace-list-start";
            const markerEnd = "replace-list-end";
            if (evt.target === null) {
                console.error("Event has no target!");
                return;
            }
            const target = evt.target;
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
            performance.measure("Typeahead list replaced", markerStart, markerEnd);
            console.log(performance.getEntriesByType("measure")[0]);
            performance.clearMarks();
            performance.clearMeasures();
        };
    }
    async function getResults(path, responseKey, body) {
        try {
            const response = await fetch(path, {
                method: "POST",
                body: JSON.stringify(body),
                headers: {
                    "Content-Type": "application/json"
                }
            });
            const jsonResponse = await response.json();
            const data = jsonResponse[responseKey];
            return data;
        }
        catch (err) {
            console.error(err);
            return [];
        }
    }
    async function vendorCallback() {
        return getResults("/vendor", "vendors", { vendor: state.vendor });
    }
    async function nameCallback() {
        return getResults("/product", "products", { vendor: state.vendor, name: state.name });
    }
    async function versionCallback() {
        return getResults("/version", "versions", { vendor: state.vendor, name: state.name, version: state.version });
    }
    wireTypeahead("vendor", "vendor-list", vendorCallback);
    wireTypeahead("name", "name-list", nameCallback);
    wireTypeahead("version", "version-list", versionCallback);
}
