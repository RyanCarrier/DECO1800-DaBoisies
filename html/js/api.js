var apikey = "j0porbqbr4efdh2c"




function get(div,zone,search){
      var URL = "http://api.trove.nla.gov.au/result?key=" + apikey + "&encoding=json&zone=" +
      zone + "&q=" + search + "&encoding=json&callback=?";
      $.getJSON(URL,function(response){
        console.log(URL);
        console.log(response);
        console.log(JSON.stringify(response));
        $(div).append(JSON.stringify(response));
      });

};



get("#search","all","britney spears");
//alert(get("all","bananas"));
//$("#search").clear();
//$('#search').append(get("all","test"));
