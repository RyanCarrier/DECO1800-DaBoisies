//run when window is loaded

var zones = ["map", "collection", "list", "people", "book", "article", "music", "picture", "newspaper"];
var names = [];
var people;
var sqbox = [0, 0, 0, 0, 0, 0];
var lastClicked;


$(window).load(function() {
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
    $("#24step4").hide();
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

jQuery(function($) {
    var open = false;
    $('#footerSlideButton').click(function() {
        if (open === false) {
            if (Modernizr.csstransitions) {
                $('#footerSlideContent').addClass('open');
            } else {
                $('#footerSlideContent').animate({
                    height: '300px'
                });
            }
            $(this).css('backgroundPosition', 'bottom left');
            open = true;
        } else {
            if (Modernizr.csstransitions) {
                $('#footerSlideContent').removeClass('open');
            } else {
                $('#footerSlideContent').animate({
                    height: '0px'
                });
            }
            $(this).css('backgroundPosition', 'top left');
            open = false;
        }
    });
});

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
    switch (tutState) {
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
            document.getElementById("nextTut").onclick = function() {
                next();
            };
            tutState = 7;
            break;
    }
}

function next() {
    switch (tutState) {
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
            document.getElementById("backTut").onclick = function() {
                back();
            };
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
            document.getElementById("backTut").onclick = function() {
                back();
            };
            document.getElementById("nextTut").onclick = function() {
                goHome2();
            };
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
    document.getElementById("backTut").onclick = function() {
        next();
    };
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
    document.getElementById("backTut").onclick = function() {
        next();
    };
    tutState = 7;
}

function goHome3() {
    $("#16veryStart").hide();
    $("#guestLeaderBoard").show();
    $("#guestPlay").show();
    $("#guestSignUp").show();
    $("#guestLogin").show();
    $("#backTut").show();
    document.getElementById("backTut").innerHTML = "<p>Try The<br>Tutorial<p>";
    document.getElementById("backTut").onclick = function() {
        next();
    };
    tutState = 0;
}

var signUpModal = document.getElementById('signUpModal');
var signUpBtn = document.getElementById("signUpBtn");
var closeSignUp = document.getElementsByClassName("closeSignUp")[0];
signUpBtn.onclick = function() {
    signUpModal.style.display = "block";
};
closeSignUp.onclick = function() {
    signUpModal.style.display = "none";
};

var loginModal = document.getElementById('loginModal');
var loginBtn = document.getElementById("loginBtn");
var closeLogin = document.getElementsByClassName("closeLogin")[0];
loginBtn.onclick = function() {
    loginModal.style.display = "block";
};
closeLogin.onclick = function() {
    loginModal.style.display = "none";
};

function playGame() {
    $("#guestLeaderBoard").hide();
    $("#guestPlay").hide();
    $("#guestSignUp").hide();
    $("#guestLogin").hide();
    $("#backTut").hide();
    $("#16veryStart").show();
    $("#gameIndicators").show();
    createCards();
}

var dmgModal = document.getElementById('dmgModal');
var dmgBtn = document.getElementById("dmgBtn");
var closeDmg = document.getElementsByClassName("closeDmg")[0];
dmgBtn.onclick = function() {
    dmgModal.style.display = "block";
};
closeDmg.onclick = function() {
    dmgModal.style.display = "none";
};

var round = 1;
var year = 2000;

document.getElementById("round").innerHTML = "Round " + round + " of 5";
document.getElementById("year").innerHTML = "Year " + year;

function toStep1() {
    $("#16veryStart").hide();
    $("#19step1").show();
}

var step1Modal = document.getElementById('step1Modal');
var closeStep1Help = document.getElementsByClassName("closeStep1Help")[0];
closeStep1Help.onclick = function() {
    step1Modal.style.display = "none";
};

function toStep2() {

    document.getElementById("s2Heading").innerHTML = "<h1>Step 2: Muzz It Out</h1>";
    document.getElementById("cont").onclick = function() {
        toStep3();
    };
    $("#cont").show();
    $("#19step1").hide();
    $("#20step2").show();
    $("#step2Modal").show();
}

var step2Modal = document.getElementById('step2Modal');
var closeStep2Help = document.getElementsByClassName("closeStep2Help")[0];
closeStep2Help.onclick = function() {
    step2Modal.style.display = "none";
};

function toStep3() {
    document.getElementById("s2Heading").innerHTML = "<h1>Step 3: Sus Out The Damage</h1>";
    document.getElementById("cont").onclick = function() {
        toStep4();
    };
    $("#23step3").show();
    $("#step2Modal").hide();
}

var step3Modal = document.getElementById('step3Modal');
var closeStep3Help = document.getElementsByClassName("closeStep3Help")[0];
closeStep3Help.onclick = function() {
    step3Modal.style.display = "none";
};

function toStep4() {
    document.getElementById("s2Heading").innerHTML = "<h1>Step 4: Mod The Squad</h1>";
    document.getElementById("cont").onclick = function() {
        toStep2();
    };
    $("#cont").hide();
    $("#23step3").hide();
    $("#24step4").show();
    $("#step4Modal").show();
    round += 1;
    year += 1;
    document.getElementById("startGame").innerHTML = "Start Next Round";
    if (round == 6) {
        $("#modSquadBtn").hide();
        document.getElementById("contSquadBtn").style.width = "20%";
        document.getElementById("contSquadBtn").style.right = "40%";
        document.getElementById("contSquadBtn").innerHTML = "Finish Game";
        document.getElementById("contSquadBtn").onclick = function() {
            finishGame();
        };
    }
}

var step4Modal = document.getElementById('step4Modal');
var closeStep4Help = document.getElementsByClassName("closeStep4Help")[0];
closeStep4Help.onclick = function() {
    step4Modal.style.display = "none";
};

function backToStep1() {
    $("#step4Modal").hide();
    $("#24step4").hide();
    toStep1();
    document.getElementById("round").innerHTML = "Round " + round + " of 5";
    document.getElementById("year").innerHTML = "Year " + year;
}

function backToStep2() {
    $("#step4Modal").hide();
    $("#20step2").hide();
    $("#24step4").hide();
    toStep2();
    document.getElementById("round").innerHTML = "Round " + round + " of 5";
    document.getElementById("year").innerHTML = "Year " + year;
}

function finishGame() {
    $("#20step2").hide();
    $("#24step4").hide();
    goHome1();
    round = 1;
    year = 2000;
    document.getElementById("round").innerHTML = "Round " + round + " of 5";
    document.getElementById("year").innerHTML = "Year " + year;
    $("#gameIndicators").hide();
}

function createCards() {
    x = 0;
    while (x < people.people.length) {
        if (people.people[x] === undefined) {
            alert(JSON.stringify(people.people[x]));
            alert(x);
        } else {
            $("#cards").append("<div " + "id=\"c" + x + "\" class=\"celeb-card col-md-1\" " + " draggable=\"true\" ondragstart=\"drag(event)\">" +
                "<img src=\"" + people.people[x].image + "\" alt=\"Mountain View\" style=\"max-height:100%;max-width:100%;\">" +
                people.people[x].query + " </div>");
        }
        x++;
    }
}

function formChecker() {
    var username = document.getElementById("username").value;
    var pass1 = document.getElementById("pass1").value;
    var pass2 = document.getElementById("pass2").value;
    if (pass1 != pass2) {
        document.getElementById("pass1").style.borderColor = "#E34234";
        document.getElementById("pass2").style.borderColor = "#E34234";
        return false;
    }
    if (username === "") {
        alert("Username has not been defined.");
        document.getElementById("username").style.borderColor = "#E34234";
        return false;
    } else if (pass1 === "") {
        alert("A password has not been set.");
        document.getElementById("pass1").style.borderColor = "#E34234";
        return false;
    } else if (pass2 !== pass1) {
        alert("Your passwords don't match.");
        document.getElementById("pass2").style.borderColor = "#E34234";
        return false;
    } else {
        alert("Passwords Match!!!");
        return true;
    }
}

function allowDrop(ev) {
    ev.preventDefault();
}

function drag(ev) {
    ev.dataTransfer.setData("text", ev.target.id);
    var id = "#" + ev.target.id;
    var parentId = $(id).parent().attr("id")
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
