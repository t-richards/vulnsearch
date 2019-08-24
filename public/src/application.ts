async function handleVendorInput(evt: Event): Promise<void> {
    if (evt.target === null) {
        return;
    }

    const target = evt.target as HTMLInputElement;
    const body = {
        vendor: target.value
    };

    const response = await fetch('/vendor', {
        method: 'POST',
        body: JSON.stringify(body),
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const jsonResponse = await response.json();
    const data = jsonResponse.vendors as Array<string>;

    const newList = document.createElement('datalist');
    newList.id = 'vendor-list';

    data.forEach((vendor) => {
        let opt = document.createElement('option');
        opt.value = vendor;
        opt.innerHTML = vendor;
        newList.appendChild(opt);
    });

    const originalList = document.getElementById('vendor-list');
    if (originalList === null || originalList.parentNode === null) {
        console.error("Failed to replace original list.");
        return;
    }
    originalList.parentNode.replaceChild(newList, originalList);
};

const vendorInput = document.getElementById('vendor');
if (vendorInput !== null) {
    vendorInput.addEventListener('input', handleVendorInput);
}
