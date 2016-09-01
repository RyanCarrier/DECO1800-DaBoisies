//Trove api key rcarrier's
var apikey = "j0porbqbr4efdh2c";
//List of the available trove zones
var zones = ["map", "collection", "list", "people", "book", "article", "music", "picture", "newspaper"];

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

function urlBuilder(zone, search) {
    search = search.replace(/ /g, "%20");
    return "http://api.trove.nla.gov.au/result?key=" + apikey + "&encoding=json&zone=" +
        zone + "&q=" + search + "&callback=?";
}

//get gets the search from the zone and appends the response to the div specified.
//In its current state it gets the relevance (with weighting) and the relevance
// without weighting.
function get(div, zone, search) {
    //TODO: Some shit when we get 500 response
    //took this out of get cause we shouldn't ever need to specify zone
    //zone = "all";
    URL = urlBuilder(zone, search);
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
    });
}

//get wrapper for the lazy
function genericGet(search) {
    get("#search", "all", search);
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
    URL = urlBuilder("list", "top") + "&include=listItems";
    console.log(URL);
    $.getJSON(URL, function(response) {
        //console.log(JSON.stringify(response.response.zone[0].records.list[0]));
        for (var item in response.response.zone[0].records.list[0].listItem) {
            item = response.response.zone[0].records.list[0].listItem[item].people;
            $("#forbes").append(JSON.stringify(item));
            $("#forbes").append("</br>");
        }
        //forbes = response;
        //$("#forbes").append(JSON.stringify(response));
    });
}

//run when window is loaded
$(window).load(function() {
    $(help).append("If nothing is coming up, check if it is 'waiting for trove' in the bottom right corner. If it is refresh the page.");
    $(help).append("Also open up the dev console for more dettails.<br><br>");
    getForbes();
    genericGet("britney spears");
    genericGet("Steve Irwin");
    genericGet("your mom");
}());
