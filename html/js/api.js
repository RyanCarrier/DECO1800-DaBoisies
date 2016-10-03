//Trove api key rcarrier's
var apikeys = [
    "j0porbqbr4efdh2c", //rcarrier
    "ulsmhsa32qhk0fhv", //robin
    "a79q82q1nosa67ck", //sam
    "8lkjcg45qi640t9s", //big dogs
    "grcr2nt2i61ourfj" //georgie
];
var keyindex = 0;
//List of the available trove zones
var zones = ["map", "collection", "list", "people", "book", "article", "music", "picture", "newspaper"];
var names = [];
var loaded = [];
var Loading = true;
//Init forbes var
var forbes;

//Create an empty object, eventually this will prbably be manually fleshed out
// when we figure out the weightings.
var ZoneWeight = {};
for (var z in zones) {
    //We will probs eventually just have this all listed out manually
    z = zones[z];
    if (z == "list") {
        ZoneWeight[z] = {
            weight: 5
        };
    } else {
        ZoneWeight[z] = {
            weight: 1
        };
    }
}
//Set up a global result var, this is more for later use when we want all the
// details saved rather than getting passed around in functions
var result;

function troveUrlBuilder(zone, search) {
    key = apikeys[keyindex];
    keyindex = (keyindex + 1) % apikeys.length;
    search = search.replace(/ /g, "%20");
    return "http://api.trove.nla.gov.au/result?key=" + key + "&encoding=json&zone=" +
        zone + "&q=" + search + "&callback=?";
}

//get gets the search from the zone and appends the response to the div specified.
//In its current state it gets the relevance (with weighting) and the relevance
// without weighting.
function get(div, zone, search, i) {
    //TODO: Some shit when we get 500 response
    //took this out of get cause we shouldn't ever need to specify zone
    //zone = "all";

    URL = troveUrlBuilder(zone, search);
    console.log(URL);
    $.getJSON(URL, function(response) {
        //Note this runs async so will need a loaded=false/true field

        //Could just pass this in, but we want to be storing this shit
        result = response;
        $(div).append("<h2>" + search + "</h2>With weighting;<br>");
        $(div).append(relevance());
        $(div).append("<br>Without weighting;<br>");
        $(div).append(relevanceNoWeighting());
        $(div).append("<br>");
        loaded[i] = true;
    });
}

//get wrapper for the lazy
function genericGet(search, i) {
    get("#search", "all", search, i);
}

function peopleUrlBuilder(search) {
    search = search.replace(/ /g, "%20");
    return "http://www.nla.gov.au/apps/srw/opensearch/peopleaustralia?q=" + search; //+ "&callback=?";
}

function relevanceNoWeighting() {
    total = 0;
    for (var z in result.response.zone) {
        //+result to force integer conversion
        total += +result.response.zone[z].records.total;
    }
    return total;
}

function relevance() {
    total = 0;
    for (var z in result.response.zone) {
        z = result.response.zone[z];
        //Don't need to + result as will convert due to multiplication
        total += z.records.total * ZoneWeight[z.name].weight;
    }
    return total;
}

function getForbes() {
    URL = troveUrlBuilder("list", "top") + "&include=listItems";
    console.log(URL);
    var nameList = [];
    idList = "";
    //  var done = false;
    $.getJSON(URL, function(response) {
        for (var item in response.response.zone[0].records.list[0].listItem) {
            item = response.response.zone[0].records.list[0].listItem[item].people[0].url;
            peopleid = JSON.stringify(item).split("/")[2].replace("\"", "");
            idList += peopleid + ",";
        }
        $.ajax({
            dataType: "text",
            url: "/api/people/" + idList.slice(0, -1),
            success: function(data) {
                names = "";
                data.split(",").forEach(function(name) {
                    names += name + ",";
                });
                nameList = names.slice(0, -1).split(",");
                for (var i = 0, len = nameList.length; i < len; i++) {
                    loaded[i] = false;
                    genericGet(nameList[i], i);
                }
                isLoaded = false;
                while (isLoaded === false) {
                    isLoaded = true;
                    /*
                                                for (i = 0, len = nameList.length; i < len; i++) {
                                                    //console.log(loaded);
                                                    if (loaded[i] === false) {
                                                        isLoaded = false;
                                                        break;
                                                    }
                        }*/
                }
                stopLoading();
            }
        });

    });
}

function stopLoading() {
    Loading = false;
    $("#floatingCirclesG").hide();
    $("#loading").hide();
}

//run when window is loaded
$(window).load(function() {
    //$(help).append("If nothing is coming up, check if it is 'waiting for trove' in the bottom right corner. If it is refresh the page.");
    //$(help).append("Also open up the dev console for more details.<br><br>");
    getForbes();
	$("#helpExit").hide();
	$("#backTut").hide();
	$("#tutBox2").hide();
	$("#tutBox3").hide();
	$("#tutBox4").hide();
	$("#tutBox5").hide();
	$("#tutBox6").hide();
	$("#tutBox7").hide();
	$("#tutBox8").hide();
	$("#tutBox9").hide();
	$("#guestLeaderBoard").hide();
	$("#guestPlay").hide();
	$("#guestSignUp").hide();
	$("#guestLogin").hide();
	$("#signUpModal").hide();
	$("#loginModal").hide();
	$("#16veryStart").hide();
	$("#dmgModal").hide();
	$("#gameIndicators").hide();
	$("#19step1").hide();
	$("#20step2").hide();
	$("#23step3").hide();
	
}());

function openHelp() {
    document.getElementById("help").style.width = "100%";
	$("#helpExit").show(1000);
}

/* Close when someone clicks on the "x" symbol inside the overlay */
function closeHelp() {
    document.getElementById("help").style.width = "0%";
	$("#helpExit").hide();
}

var tutState = 0;
function back() {
	switch(tutState) {
		case 1:
			$("#tutBox").show();
			$("#skipButton").show();
			$("#backTut").hide();
			$("#tutBox2").hide();
			document.getElementById("nextTut").innerHTML = "<p>Start<br>Tutorial<p>";
			tutState = 0;
			break;
		case 2:
			$("#tutBox2").show();
			$("#tutBox3").hide();
			tutState = 1;
			break;
		case 3:
			$("#tutBox3").show();
			$("#tutBox4").hide();
			tutState = 2;
			break;
		case 4:
			$("#tutBox4").show();
			$("#tutBox5").hide();
			tutState = 3;
			break;
		case 5:
			$("#tutBox5").show();
			$("#tutBox6").hide();
			tutState = 4;
			break;
		case 6:
			$("#tutBox6").show();
			$("#tutBox7").hide();
			tutState = 5;
			break;
		case 7:
			$("#tutBox7").show();
			$("#tutBox8").hide();
			tutState = 6;
			break;
		case 8:
			$("#tutBox8").show();
			$("#tutBox9").hide();
			document.getElementById("nextTut").innerHTML = "<p>Next<p>";
			document.getElementById("nextTut").onclick = function (){ next() };
			tutState = 7;
			break;
	}
}

function next() {
	switch(tutState) {
		case 0:
			$("#tutBox").hide();
			$("#skipButton").hide();
			$("#backTut").show();
			$("#tutBox2").show();
			$("#guestLeaderBoard").hide();
			$("#guestPlay").hide();
			$("#guestSignUp").hide();
			$("#guestLogin").hide();
			$("#nextTut").show();
			document.getElementById("nextTut").innerHTML = "<p>Next<p>";
			document.getElementById("backTut").innerHTML = "<p>Back<p>";
			document.getElementById("backTut").onclick = function (){ back() };
			tutState = 1;
			break;
		case 1:
			$("#tutBox2").hide();
			$("#tutBox3").show();
			tutState = 2;
			break;
		case 2:
			$("#tutBox3").hide();
			$("#tutBox4").show();
			tutState = 3;
			break;
		case 3:
			$("#tutBox4").hide();
			$("#tutBox5").show();
			tutState = 4;
			break;
		case 4:
			$("#tutBox5").hide();
			$("#tutBox6").show();
			tutState = 5;
			break;
		case 5:
			$("#tutBox6").hide();
			$("#tutBox7").show();
			tutState = 6;
			break;
		case 6:
			$("#tutBox7").hide();
			$("#tutBox8").show();
			tutState = 7;
			break;
		case 7:
			$("#tutBox8").hide();
			$("#tutBox9").show();
			$("#guestLeaderBoard").hide();
			$("#guestPlay").hide();
			$("#guestSignUp").hide();
			$("#guestLogin").hide();
			$("#nextTut").show();
			document.getElementById("nextTut").innerHTML = "<p>Finish<br>Tutorial<p>";
			document.getElementById("backTut").innerHTML = "<p>Back<p>";
			document.getElementById("backTut").onclick = function (){ back() };
			document.getElementById("nextTut").onclick = function (){ goHome2() };
			tutState = 8;
			break;
	}
}

function goHome1() {
	$("#tutBox").hide();
	$("#skipButton").hide();
	$("#tutBox9").hide();
	$("#nextTut").hide();
	$("#guestLeaderBoard").show();
	$("#guestPlay").show();
	$("#guestSignUp").show();
	$("#guestLogin").show();
	$("#backTut").show();
	document.getElementById("backTut").innerHTML = "<p>Try The<br>Tutorial<p>";
	document.getElementById("backTut").onclick = function (){ next() };
	tutState = 0;
}


function goHome2() {
	$("#tutBox").hide();
	$("#skipButton").hide();
	$("#tutBox9").hide();
	$("#nextTut").hide();
	$("#guestLeaderBoard").show();
	$("#guestPlay").show();
	$("#guestSignUp").show();
	$("#guestLogin").show();
	$("#backTut").show();
	document.getElementById("backTut").innerHTML = "<p>Back To<br>Tutorial<p>";
	document.getElementById("backTut").onclick = function (){ next() };
	tutState = 7;
}

var signUpModal = document.getElementById('signUpModal');
var signUpBtn = document.getElementById("signUpBtn");
var closeSignUp = document.getElementsByClassName("closeSignUp")[0];
signUpBtn.onclick = function() {
    signUpModal.style.display = "block";
}
closeSignUp.onclick = function() {
    signUpModal.style.display = "none";
}

var loginModal = document.getElementById('loginModal');
var loginBtn = document.getElementById("loginBtn");
var closeLogin = document.getElementsByClassName("closeLogin")[0];
loginBtn.onclick = function() {
    loginModal.style.display = "block";
}
closeLogin.onclick = function() {
    loginModal.style.display = "none";
}

function playGame() {
	$("#guestLeaderBoard").hide();
	$("#guestPlay").hide();
	$("#guestSignUp").hide();
	$("#guestLogin").hide();
	$("#backTut").hide();
	$("#16veryStart").show();
	$("#gameIndicators").show();
}

var dmgModal = document.getElementById('dmgModal');
var dmgBtn = document.getElementById("dmgBtn");
var closeDmg = document.getElementsByClassName("closeDmg")[0];
dmgBtn.onclick = function() {
    dmgModal.style.display = "block";
}
closeDmg.onclick = function() {
    dmgModal.style.display = "none";
}

function toStep1() {
	$("#16veryStart").hide();
	$("#19step1").show();
}

var step1Modal = document.getElementById('step1Modal');
var closeStep1Help = document.getElementsByClassName("closeStep1Help")[0];
closeStep1Help.onclick = function() {
    step1Modal.style.display = "none";
}

function toStep2() {
	$("#19step1").hide();
	$("#20step2").show();
}

function toStep3() {
	document.getElementById("s2Heading").innerHTML = "<h1>Step 3: Sus Out The Damage</h1>";
	document.getElementById("cont").onclick = function (){ toStep4() };
	$("#23step3").show();
}


var step3Modal = document.getElementById('step3Modal');
var closeStep3Help = document.getElementsByClassName("closeStep3Help")[0];
closeStep3Help.onclick = function() {
    step3Modal.style.display = "none";
}