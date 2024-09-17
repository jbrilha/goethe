/**
 * @param {HTMLInputElement} inputTag
 * @param {HTMLInputElement} formTags
 */
function addTag(inputTag, formTags) {
    const tag = removeWhiteSpace(inputTag.value);

    if (!validateTag(tag, formTags.value)) { inputTag.value = tag; return; }

    var tags = document.getElementById("tags");
    const tagElem = tags.querySelector("#tag");

    var clone = tagElem.cloneNode(true);

    // @ts-ignore
    clone.classList.remove("hidden")
    // @ts-ignore
    var tagName = clone.querySelector("#tag-name")
    tagName.innerHTML = tag

    tags.appendChild(clone)

    _hyperscript.processNode(clone)

    inputTag.value = "";
    formTags.value += tag + " ";
}

/**
 * @param {string} tag
 * @param {string} currentTags
 * @returns {boolean}
 */
function validateTag(tag, currentTags) {
     if (tag === "") return false;

    if (tag.match(RegExp('^[a-zA-Z0-9_]+$')) == null) {
        showNotif("Tags must be alphanumeric with no spaces.<br>_underscores_ are okay ;D")
        return false
    }

    var tagsArray = currentTags.split(" ").map((tag) => tag.trim());

    if (tagsArray.includes(tag)) {
        showNotif("Tag '" + tag + "' already added!")
        return false;
    }

    
    return true
}

/**
 * @param {string} msg
 */
function showNotif(msg) {
    const notif = `
        <div
            id="alert"
            class="m-1 bg-white border-black border-2 p-2 text-lg w-fit max-w-xs flex items-stretch justify-between fade"
            _="on load wait 5s add .fade-hidden to me wait 0.4s remove me"
        >
            <p id="alert-msg" class="pr-2 font-bold">
                ${msg}
            </p>
            <button
                id="close-button"
                class="h-full flex-grow-0 px-2 btn-black"
                type="button"
                _="on click 
                    add .fade-hidden to my.parentElement
                    wait 0.4s 
                    remove my.parentElement"
            >
                &times;
            </button>
        </div>
    `

    var notifs = document.querySelector("#notifications")
    const alertDiv = document.createElement('div');
    alertDiv.innerHTML = notif;

    _hyperscript.processNode(alertDiv);

    notifs.appendChild(alertDiv);
}


/**
 * @param {string} tag2remove
 * @param {HTMLInputElement} formTags
 */
function removeTag(tag2remove, formTags) {
    var tagsArray = formTags.value.split(",").map((tag) => tag.trim());

    var filteredTags = tagsArray.filter((tag) => tag != tag2remove);

    formTags.value = filteredTags.join(",");

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
var htmx, _hyperscript;
