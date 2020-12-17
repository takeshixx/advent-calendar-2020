"strict";

const listWrapper = document.getElementById("list-wrapper");
const template = document.getElementById("naughty-list-tmpl").innerHTML;
const view = {
    "people": JSON.parse(listWrapper.getAttribute("data-list")),
    "name": function () {
      return this.firstName + " " + this.lastName;
    },
    "index": function() {
        return ++window['INDEX']||(window['INDEX']=0);
    }
  };
const input = document.querySelector("#add");

function renderNaughtyList() {
    const rendered = Mustache.render(template, view);
    listWrapper.innerHTML = rendered;
    updateEventListener();
}

renderNaughtyList();

window.deleteEntry = function(index) {
    view.people.splice(index, 1);
    renderNaughtyList();
};

function updateEventListener() {
    document
    .querySelectorAll("a[data-action]")
    .forEach(n => {
        n.addEventListener("click", e => {
            const action = JSON.parse(e.target.getAttribute("data-action"));
            console.log(action);
            globalThis[action.target ?? "window"][action.method](...(action.parameters ?? [e.target]));
            return false;
        });
    });
}

document.querySelector("form").addEventListener("submit",(e) => {
    e.preventDefault();
    const person = input.value.split(" ", 2);
    if(person.length != 2) {
        alert("Invalid name format, expect: firstName lastName");
        return;
    }
    input.value = "";
    view.people.push({ "firstName": person[0], "lastName": person[1] });
    renderNaughtyList();
});
