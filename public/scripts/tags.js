/**
 * @param {string} tag
 */
function addTag(tag) {
  if (tag === "") return;

  var inputTags = document.getElementById("input-tags");

  if (!inputTags || !(inputTags instanceof HTMLInputElement)) {
    console.error("Couldn't find #input-tags");
    return;
  }

  var currentTags = inputTags.value.trim();
  var tagsArray = currentTags.split(",").map((tag) => tag.trim());

  if (tagsArray.includes(tag)) {
    const alert = encodeURIComponent("Tag '" + tag + "' already added!");
    htmx.ajax("GET", `/alert?a=${alert}`);
    return;
  }

  const path = `/posts/create?tag=${encodeURIComponent(tag)}`;
  const ctx = { target: "#tags", swap: "beforeend" };

  htmx.ajax("PUT", path, ctx).then(() => {
    var tags = document.getElementById("tags");

    if (!inputTags || !(inputTags instanceof HTMLInputElement)) {
      console.error("Couldn't find #input-tags");
      return;
    }

    if (currentTags) {
      inputTags.value = currentTags + "," + tag;
    } else {
      inputTags.value = tag;
    }
    if (tags.children.length == 1) {
      tags.classList.add("mb-4");
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

// need this to avoid tsserver errors
/**
 * @type {any}
 */
var htmx;
