function allowDrop(ev) {
    ev.preventDefault();
}

function drag(ev) {
    ev.dataTransfer.setData("text", ev.target.id);
    var id = "#" + ev.target.id;
    var parentId = $(id).parent().attr("id");
    lastClicked = parentId;
}

function dropDefault(ev) {
    //$("#cards").append(ev.);
    //ev.target.append(document.getElementById("cards"));
    var data = ev.dataTransfer.getData("text");
    ev.target.appendChild(document.getElementById(data));

}

function drop(ev) {
    ev.preventDefault();
    var target = $(event.target);
    console.log(target.attr('id'));
    var id = target.attr("id");
    var i = +id[2] - 1;
    if (id.length == 3 && id[0] == 's' && id[1] == 'm') {
        if (sqbox[i] === 0) {
            dropDefault(ev);
            sqbox[i] = 1;
            id = lastClicked;
            a = +id[2] - 1;
            sqbox[a] = 0;
        } else {
            console.log("nup");
        }
    } else if (id === "cards") {
        dropDefault(ev);
        id = lastClicked;
        i = +id[2] - 1;
        sqbox[i] = 0;
        console.log(i);
    } else if (id[0] == 'c') {
        console.log("don't put me on another card u retard");
    } else {
        dropDefault(ev);
    }
}
