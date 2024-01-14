$(document).ready(() => {
    deleteNewsById = (id) => {
        const deleteUrl = "/news/delete?id=" + id;

        $.ajax({
            url: deleteUrl,
            method: 'DElETE',
            success: function(response) {
                console.log('DELETE request successful', response);
                location.reload()
            },
            error: function(error) {
                console.error('Error in DELETE request', error);
            }
        });
    }
})

$(document).ready(() => {
    UpdateNewsById=(id) => {
        const deleteUrl = "/news/delete?id=" + id;

        const title = $("#update-title").val();
        const content = $("#update-content").val();
        const category = $("#update-category").val();


        $.ajax({
            url: updateUrl,
            method: 'POST', // Assuming your server expects a POST request for updates
            data: {
                title: title,
                content: content,
                category: category
            },
            success: function (response) {
                console.log('Update request successful', response);
                location.reload();
            },
            error: function (error) {
                console.error('Error in update request', error);
            }
        });
    }

    // Call updateNews function when the document is ready or when the update button is clicked
    $("#update-button").onclick,UpdateNewsById;
});
