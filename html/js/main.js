var state = 0;
var year = 2000;
var round = 1;
var totalScore = 0;
//0 main menu
//1 submain menu
//2 choose squad
//3 assess damage(do you want to edit or continue) -> loop to 2
//4 final score

/*
"How does it work?",
"Gatho The Squad",
"Sus Out The Damage",
"Mod Your Squad",
"End game"*/

var zones = ["map", "collection", "list", "people", "book", "article", "music", "picture", "newspaper"];
var names = [];
var people;
var sqbox = [0, 0, 0, 0, 0, 0];
var lastClicked;

function help() {
    helpModal(state);
}

function next() {
    switch (state) {
        case 0:
            $('#homeOptions').show();
            $('#home').hide();
            //I want the logo tos tart big and shrink small
            break;
        case 1:
            $('#homeOptions').hide();
            $('#squad').show();
            $('#nextBtn').html("Submit Squad");
            doGameIndicator();
            $('#gameIndicators').show();
            createCards();
            break;
        case 2:
            if (!(isSquadFull())) {
                return;
            }
            runDamage();
            $("#squadSummary").show();
            $("#cards").hide();
            break;
        case 3:
            round++;
            year++;
            $("#squadSummary").show();
            $("#cards").hide();
            if (round > 5) {
                state--;
                next();
                return;
            }
            $("#squadSummary").html("Final score is " + totalScore);
    }
    state++;
}

function runDamage() {
    //get names from thingy
    //loop over them
    //print to modal with spinner while loading

}

function isSquadFull() {
    for (var i = 1; i <= 6; i++) {
        if ($('#sm' + i).is(':empty')) {
            alert("Please fill slot " + i + " with a celeb!");
            return false;
        }
    }
    return true;
}

function doGameIndicator() {
    $('#roundIndicator').html("<p>Round " + round +
        " of 5.</p> <p> Current year: " + year + " </p>" +
        "<p><b>Current score: " + totalScore + "</b></p>");
}

function home() {
    $('#homeOptions').hide();
    $('#squad').hide();
    $('#gameIndicators').hide();
    $('#home').show();
    $("#squadSummary").hide();
    $("#squadSummary").html("We are sorry but no squad summary is available at this time.");
    $("#cards").show();
    state = 0;
    round = 1;
    year = 2000;
    totalScore = 0;
}


function createCards() {
    x = 0;
    while (x < people.people.length) {
        if (people.people[x] === undefined) {
            alert(JSON.stringify(people.people[x]));
            alert(x);
        } else {
            $("#cards").append("<div " + "id=\"c" + x + "\" class=\"celeb-card col-md-1\" " + " draggable=\"true\" ondragstart=\"drag(event)\">" +
                "<img src=\"" + people.people[x].image + "\" style=\"max-height:100%;max-width:100%;\">" +
                people.people[x].query + " </div>");
        }
        x++;
    }
}

$(window).load(function() {
    $.ajax({
        type: 'GET',
        url: "/api/getlist/",
        async: false,
        contentType: "application/json",
        dataType: 'json',
        success: function(json) {
            people = json;
        },
        error: function(e) {
            alert("error getting data");
        }
    });
    //$('#playBtn').removeClass("hidden");
    //$('#loginBtn').removeClass("hidden");
    //$('#leaderboardBtn').removeClass("hidden");
    //setupListeners();
});
