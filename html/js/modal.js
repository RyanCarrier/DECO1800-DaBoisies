/*jshint esversion: 6 */
function createModal(heading, body) {
    setModalHead(heading);
    setModalBody(body);
    showModal();
}


function setModalHead(header) {
    $("#myModalHead").html(header);
}

function setModalBody(body) {
    $("#myModalBody").html(body);
}

function hideModal() {
    $("#myModal").modal('hide');
}

function showModal() {
    $("#myModal").modal('show');
}

function sorry() {
    createModal(':\'(', 'Sorry... This feature is not implemented yet...');
}

function showLeaderboard() {
    body = `<div class="table table-responsive container-fluid center-block">\
    <table id="board" class="table table-responsive container-fluid center-block">\
        <tr>
            <th>Ranking</th>
            <th>Username</th>
            <th>Party Score</th>
        </tr>
        <tr>
            <td>#1</td>
            <td>carrierpigeon</td>
            <td>68189</td>
        </tr>
        <tr>
            <td>#2</td>
            <td>iamrobin</td>
            <td>59894</td>
        </tr>
        <tr>
            <td>#3</td>
            <td>anotherbirdthing</td>
            <td>32654</td>
        </tr>
    </table>
    <!-- show user's game history -->
</div>`;
    head = "<h2>Leaderboards</h2>";
    createModal(head, body);

}
