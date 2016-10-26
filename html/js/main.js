/*jshint esversion: 6 */
var state = 0;
var year = 2000;
var round = 1;
var totalScore = 0;
var squadLoaded = [false, false, false, false, false, false];
var squadScore = [0, 0, 0, 0, 0, 0];
var roundScore = 0;
/*STATE SUMMARY;

0 main menu
1 submain menu
2 choose squad
3 assess damage -> loop to 2 (four times)
4 final score
*/

var zones = ["map", "collection", "list", "people", "book", "article", "music", "picture", "newspaper"];
var names = [];
var people;


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
            $('#nextBtn').html("Next");
            break;
        case 3:
            round++;
            year++;
            $("#squadSummary").hide();
            $("#cards").show();
            $('#nextBtn').html("Submit Squad");
            $("#game-head").html("Mod your Squad");
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
    $("#theSquad").children("div").each(function(index, element) {
        $(element).each(function(i, e) {
            individualRunDamage(i, e);
        });
        //individualRunDamage(index, element);
    });
    for (var i = 0; i <= squadLoaded.length; i++) {
        if (squadLoaded[i] === false) {
            // Usage!
            //alert(JSON.stringify(squadLoaded));
            //  sleep(500).then(() => {});
            i = -1;
        }
    }
    createModal("Damage report", "You scored " + roundScore + " this round with " + "someone" + " scoring the most points for you!");

    //get names from thingy
    //loop over them
    //print to modal with spinner while loading

}

function sleep(time) {
    return new Promise((resolve) => setTimeout(resolve, time));
}

function individualRunDamage(index, element) {
    name = $(element).text().trim();
    alert("/api/weight/" + encodeURI(name) + "/year/" + year);
    $.getJSON("/api/weight/" + encodeURI(name) + "/year/" + year, function(got) {
        total = 0;
        for (var i = 0; i < got.zones.length; i++) {
            total += got.zones[i].total;
        }
        total *= 47;
        total /= 7;
        alert(index);
        squadScore[index] = total;
        squadLoaded[index] = true;
    });
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
    $("#game-head").html("Choose your Squad");
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
                "<img src=\"" + people.people[x].image + "\" style=\"max-height:100%;max-width:100%;\"><p>" +
                people.people[x].query + " </p></div>");
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
});
