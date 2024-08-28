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

// I do ned to get the tag partial from the server otherwise _hyperscrit won't
// register the actions and therefore each tag becomes impossible to remove
//
// TODO: eventually add event listeners to the tag to avoid these server calls

// /**
//  * @param {string} tag
//  */
// function addTagClientSide(tag) {
//   console.log(tag);
//   if (tag === "") return;
//
//   var inputTags = document.getElementById("input-tags");
//
//   if (!inputTags || !(inputTags instanceof HTMLInputElement)) {
//     console.error("Couldn't find #input-tags");
//     return;
//   }
//
//   var currentTags = inputTags.value.trim();
//   var tagsArray = currentTags.split(",").map((tag) => tag.trim());
//   console.log("currtags", currentTags);
//
//   if (tagsArray.includes(tag)) {
//     console.warn("tag already present");
//     // const alert = encodeURIComponent("Tag '" + tag + "' already added!");
//     // htmx.ajax("GET", `/alert?a=${alert}`);
//     return;
//   }
//
//   let tagElement = document.createElement("div");
//   tagElement.id = "tag";
//   tagElement.classList.add(
//     "flex",
//     "items-stretch",
//     "rounded-full",
//     "bg-yellow-200",
//     "px-1",
//     "m-1",
//   );
//   tagElement.innerHTML = `
//         <div class="px-2 flex items-center">
// 			<span>#</span>
// 			<p id="tag-name">
// 				${tag}
// 			</p>
// 		</div>
// 		<button
// 			id="remove-tag"
// 			class="h-full flex-grow-0 px-2 border-l border-black hover:font-bold"
// 			type="button"
// 			_="on click
//                 call removeTag((previous <p#tag-name/>).innerText)
//                 then remove the closest #tag"
// 		>
// 			&times;
// 		</button>
// `;
//
//   var tags = document.getElementById("tags");
//   tags.appendChild(tagElement);
//
//   if (currentTags) {
//     inputTags.value = currentTags + "," + tag;
//   } else {
//     inputTags.value = tag;
//   }
//
//   if (tags.children.length == 1) {
//     tags.classList.add("mb-4");
//   }
// }

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

// need this to avoid tsserver errors
/**
 * @type {any}
 */
var htmx;
