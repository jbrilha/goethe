/**
 * @param {string} tag
 */
function addTag(tag) {
    if (tag === "") return;

    // TODO maybe it would be smarter to switch to regular browser alerts
    if (tag.includes(' ')) {
        const alertMsg = encodeURIComponent("No spaces allowed!");
        htmx.ajax("GET", `/alert?a=${alertMsg}`);
        return;
    }

    var inputTags = document.getElementById("input-tags");

    if (!inputTags || !(inputTags instanceof HTMLInputElement)) {
        console.error("Couldn't find #input-tags");
        return;
    }

    var currentTags = inputTags.value.trim();
    var tagsArray = currentTags.split(",").map((tag) => tag.trim());

    if (tagsArray.includes(tag)) {
        const alertMsg = encodeURIComponent("Tag '" + tag + "' already added!");
        htmx.ajax("GET", `/alert?a=${alertMsg}`);
        return;
    }

    const path = `/posts/create?tag=${encodeURIComponent(tag)}`;
    const ctx = { target: "#tags", swap: "beforeend" };

    var tags = document.getElementById("tags");
    const tagsLen = tags.children.length;

    htmx.ajax("PUT", path, ctx).then(() => {
        if (!inputTags || !(inputTags instanceof HTMLInputElement)) {
            console.error("Couldn't find #input-tags");
            return;
        }

        if (currentTags) {
            inputTags.value = currentTags + "," + tag;
        } else {
            inputTags.value = tag;
        }
        if (tags.children.length != tagsLen) {
            const tagInput = document.getElementById("tag-input");
            if (!tagInput || !(tagInput instanceof HTMLInputElement)) {
                return;
            }
            tagInput.value = "";
        }
    });
}

/**
 * @param {string} tag2remove
 */
function removeTag(tag2remove) {
    var inputTags = document.getElementById("input-tags");

    if (!inputTags || !(inputTags instanceof HTMLInputElement)) {
        console.error("Couldn't find #input-tags");
        return;
    }

    var tagsArray = inputTags.value.split(",").map((tag) => tag.trim());

    var filteredTags = tagsArray.filter((tag) => tag != tag2remove);

    inputTags.value = filteredTags.join(",");

    var tags = document.getElementById("tags");
    if (tags.children.length == 1) {
        tags.classList.remove("mb-4");
    }
}

/**
 * @param {string} str
 * @returns {string}
 */
function removeWhiteSpace(str) {
    return str.replace(/\s+/g, " ").trim();
}

// need this to avoid tsserver errors
/**
 * @type {any}
 */
var htmx;
