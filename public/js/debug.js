var sidebar = false;
$('#toggleSidebar').on('click', function() {
    if (sidebar == true) {
        $('#sidebar').toggle("slide");
        sidebar = !sidebar;
    } else {
        $('#sidebar').toggle("slide");
        sidebar = !sidebar;
    }
});
