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
    r += '        <p>Image Title</p>';
    r += '        <input id="image_title" placeholder="A day at the beach" name="frame.' + n + '.title" value="" type="text">';
    r += '    </div>';

    r += '    <div>';
    r += '        <p for="image_description">Image Description</p>';
    r += '        <input id="image_description" rows="3" cols="25" name="frame.' + n + '.description">';
    r += '    </div>';

    r += '    <div>';
    r += '        <p for="image_link">Image Link</p>';
    r += '        <input id="image_link" name="frame.' + n + '.link" type="text">';
    r += '    </div>';

    r += '    <div>';
    r += '        <input type="button" id="delete_frame" value="Delete Frame">';
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
