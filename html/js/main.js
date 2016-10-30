/*jshint esversion: 6 */
var state = 0;
var year = 2000;
//from 1980
//max 2014
var round = 1;
var totalScore = 0;
var squadLoaded = [false, false, false, false, false, false];
var squadScore = [0, 0, 0, 0, 0, 0];
var squad = ["", "", "", "", "", ""];
var roundScore = 0;
var xx = 0;
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
            minLogo();
            //I want the logo to start big and shrink small
            break;
        case 1:
            $('#homeOptions').hide();
            $('#squad').show();
            $('#nextBtn').html("Submit Squad");
            doGameIndicator();
            $('#gameIndicators').show();
            break;
        case 2:
            if (!(isSquadFull())) {
                return;
            }
            $("#squadSummary").show();
            $("#cards").hide();
            if (round < 4) {
                $("#nextBtn").html(`Mod my squad!`);
            } else {
                $("#nextBtn").html(`Final summary!`);
            }
            runDamage();
            round++;
            year++;
            doGameIndicator();
            //$('#nextBtn').html("Next");
            break;
        case 3:
            $("#squadSummary").hide();
            $("#cards").show();
            $('#nextBtn').html("Submit Squad");
            $("#game-head").html("Mod your Squad");
            if (round < 5) {
                squadLoaded = [false, false, false, false, false, false];
                squadScore = [0, 0, 0, 0, 0, 0];
                squad = ["", "", "", "", "", ""];
                state -= 2;
                next();
                return;
            }
            createModal("Finay away!", "Your final score was " + totalScore + " great job!! ish...");
            //$("#squadSummary").html("Final score is " + totalScore);
            home();
    }
    state++;
}

function runDamage() {
    $("#theSquad").children("div").each(function(index, element) {
        $(element).each(function(i, e) {
            console.log(JSON.stringify(squad));
            //alert(xx);
            squad[xx] = $(element).text().trim();
            individualRunDamage(i, e);
            xx++;
        });
        //xx = 0;
        //individualRunDamage(index, element);
    });
    createModal("Muzzing...", `<span class="glyphicon glyphicon-refresh glyphicon-refresh-animate"></span>`);
    setTimeout(hasLoaded, 100);



    //get names from thingy
    //loop over them
    //print to modal with spinner while loading

}

function hasLoaded() {
    for (var i = 0; i <= squadLoaded.length; i++) {
        if (squadLoaded[i] === false) {
            setTimeout(hasLoaded, 100);
            return;
        }
    }
    //hideModal();
    console.log("loaded...");
    createModal("Damage report", "You scored " + roundScore + " this round with " + "someone" + " scoring the most points for you!");
    totalScore += roundScore;

}

function sleep(time) {
    return new Promise((resolve) => setTimeout(resolve, time));
}

function individualRunDamage(index, element) {
    name = $(element).text().trim();
    name = name.replace("/", "%2F");
    //alert("/api/weight/" + encodeURI(name) + "/year/" + year);
    $.getJSON("/api/weight/" + encodeURI(name) + "/year/" + year, function(got) {
        total = 0;
        for (var i = 0; i < got.zones.length; i++) {
            total += got.zones[i].total;
        }
        total *= 47;
        total /= 7;
        total = Math.round(total);
        roundScore += total;
        console.log(total);
        for (i = 0; i < squad.length; i++) {
            if (squad[i] === got.query) {
                squadScore[i] = total;
                squadLoaded[i] = true;
                console.log("" + i + " done " + got.query);
                return;
            }
        }
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
    maxLogo();
    state = 0;
    round = 1;
    year = 2000;
    totalScore = 0;
}


function createCards() {
    console.log("createcards");
    x = 0;
    while (x < people.people.length) {
        if (people.people[x] === undefined) {
            alert(JSON.stringify(people.people[x]));
            alert(x);
        } else {
            $("#cards").append("<div " + "id=\"c" + x + "\" class=\"celeb-card\" " + " draggable=\"true\" ondragstart=\"drag(event)\">" +
                "<img src=\"" + people.people[x].image + "\" style=\"max-height:70px;max-width:80px;pointer-events:none;\"><p>" +
                people.people[x].query + " </p></div>");
        }
        x++;
    }
}

function minLogo() {
    var img = $("#logoID");
    img.animate({
        height: "15%",
        width: "17%"
    }, 1000);
}

function maxLogo() {
    var img = $("#logoID");

    img.animate({
        height: "75%",
        width: "78%"
    }, 1000);
}
$(window).load(function() {
    s = "";
    for (var i = 1; i <= 6; i++) {
        s += "<div id=sq" + i + "></div>";
    }
    maxLogo();
    $("#datacontainer").html(s);
    //alert("trying to get");
    $.ajax({
        type: 'GET',
        url: "/api/getlist/",
        async: false,
        contentType: "application/json",
        dataType: 'json',
        success: function(json) {
            people = json;
            createCards();
        },
        error: function(e) {
            alert("error getting data" + JSON.stringify(e));
            alert(e);

        }
    });
});
