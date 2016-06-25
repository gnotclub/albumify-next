(function() {

frames = 0;

function updateNumbers() {
    $('.number').each(function(i, el) {
        el.innerHTML = 'Image ' + (i + 1);
    });
};

function addFrame(n) {
    frames++;
    r  = '<div id="frame' + n + '" class="frame">';
    r += '    <h3 class="number">Image ' + (n + 1) + '</h1>';
    r += '    <div>';
    r += '        <p class="label" for="image_title">Image Title</p>';
    r += '        <input id="image_title" placeholder="A day at the beach" name="frame.' + n + '.title" value="" type="text">';
    r += '    </div>';

    r += '    <div>';
    r += '        <p class="label" for="image_description">Image Description</p>';
    r += '        <textarea id="image_description" rows="2" cols="25" name="frame.' + n + '.description" placeholder="37 degrees outside"></textarea>';
    r += '    </div>';

    r += '    <div>';
    r += '        <p class="label" for="image_link">Image Link</p>';
    r += '        <input id="image_link" name="frame.' + n + '.link" type="text" placeholder="http://i.imgur.com/turtles.png">';
    r += '    </div>';

    r += '    <div>';
    r += '        <button id="delete_frame">Delete Frame</button>';
    r += '    </div>';
    r += '</div>';

    $('#frames').append(r);
    updateNumbers();
    $('#frame' + n + ' #delete_frame').on('click', function() {
        $('#frame' + n).remove();
        updateNumbers();
    });
}

addFrame(frames);
$('#addFrame').on('click', function () { addFrame(frames); });

})();
