$(document).ready(() => {
    deleteNewsById = (id) => {
        const deleteUrl = "/news/delete?id=" + id;

        $.ajax({
            url: deleteUrl,
            method: 'DELETE',
            success: function (response) {
                console.log('DELETE request successful', response);
                location.reload();
            },
            error: function (error) {
                console.error('Error in DELETE request', error);
            }
        });
    }
});
