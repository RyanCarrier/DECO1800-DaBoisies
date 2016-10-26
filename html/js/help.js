title = "Welcome to Ultimate Squad Goals Simulator 2K16";
instructions = "Ultimate Squad Goals Simulator 2K16 is a fantasy party simulator where you must assemble a hypothetical squad of real life celebrities in an attempt to host the best parties ever. Click on the arrow to your right to begin the tutorial.";
objective = "The objective of the game is to host the best parties for five consecutive years in a row by inviting the most popular celebrities of those years.";
howitworks = "When you start a new game, you are provided with a random year between 2000 and 2010. This is the year you will begin your 5 consecutive yearly parties. You will loop through step 1 to step 4 for five consecutive years in a row. Your party score for each year is totalled to calculate your total party score. Depending on whether or not you want your score to be recorded, you can add you score to the leader board once you've finished playing for the five consecutive years.";
leaderboards = "You can play as a guest or you can login and create an account. The benefit of creating an account is that when you add your total party score to the leader board, your username will be posted as well - letting everyone know who's boss.";

helpArray = [
    howitworks,
    howitworks,
    "Your first step is to invite celebrities to your party - these celebrities are your squad members. Who you choose to invite to your squad will directly influence how successful your party is; the most popular celebrities (of each year) will better influence your party score. When you finish building your squad, submit your squad to finalize the invitations.",
    "After submitting your squad, your party score is evaluated based on who you invited to be in your squad. The scoring system for each celebrity is calculated based on their time in the spotlight over the years. Each celebrity is given a popularity rating each year, which can increase or decrease depending on their career choices. Your damage report contains your party score for the year and individual assessment reports for each of your squad memers detailing: if they positively/negatively influenced the party score and what they were doing during that year (which may explain their influence on your party). You have now passed your first year of your 5 consecutive yearly partying and have been provided with a damage report. Given the success or failure of your party, you start planning for next year's party by modifying your squad. You can continue with your current squad if you believe they will do well next year, or you can modify your squad to switch things up.",
    "You may now choose to play again to try beat your old score! In the future you will be able to save your score to our leaderboard and verse friends live!"
];

helpTitles = [
    "How does it work?",
    "How does it work?",
    "Gatho The Squad",
    "Sus Out The Damage",
    "End game"
];

function getHelp(index) {
    if (index < 0 || index > 3) {
        return "<h2> ur coder is retarded </h2>";
    }
    return "<h2>" + helpTitles[index] + "</h2><p>" + helpArray[index] + "</p>";
}

function getHelpAll() {
    retVal = "<h1>Help!</h1>" + "<br>";
    retVal += "<h2>How it works!</h2><p>" + howitworks + "</p><br>";
    for (var i = 0; i <= 3; i++) {
        retVal += getHelp(i) + "<br>";
    }
    retVal += "<h2>Leaderboards? </h2><p>" + leaderboards + "</p><br>";
}

function helpModal(index) {
    createModal("<h2>" + helpTitles[index] + "</h2>", "<p>" + helpArray[index] + "</p>");
}
