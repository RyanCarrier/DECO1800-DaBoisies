var apikey = "j0porbqbr4efdh2c";

var zones = ["map", "collection", "list", "people", "book", "article", "music", "picture", "newspaper"];
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
var result;

function get(div, zone, search) {
    //TODO: Some shit when we get 500 response
    var URL = "http://api.trove.nla.gov.au/result?key=" + apikey + "&encoding=json&zone=" +
        zone + "&q=" + search + "&callback=?";
    $.getJSON(URL, function(response) {
        //Note this runs async so will need a loaded=false/true field
        console.log(URL);
        //Could just pass this in, but we want to be storing this shit
        result = response;
        $(div).append(relevance());
        $(div).append("<br>");
        $(div).append(relevanceNoWeighting());
    });
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




get("#search", "all", "britney spears");
//alert(get("all","bananas"));
//$("#search").clear();
//$('#search').append(get("all","test"));
