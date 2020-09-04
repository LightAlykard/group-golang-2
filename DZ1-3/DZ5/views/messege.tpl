{{template "UIkit"}}
{{template "JS"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>BeeGo posts!</title>
</head>
<body>
    <div messege-id="{{.Messeges.Id}}" class="uk-card uk-card-default uk-card-body">
<h1>{{.Messeges.Id}})<input type="text" class="uk-input" name="Name" value="{{.Messeges.Name}}"/> </h1>
    <hr>
    <input type="text" class="uk-input" name="Text" value="{{.Messeges.Text}}"/> 
    <br>
    <hr>
</div>
        <button class="uk-button uk-button-default" onclick="updateMessege('{{.Messeges.Id}}')">Edit</button>
        <button class="uk-button uk-button-default" onclick="deleteMessege('{{.Messeges.Id}}')">Delete</button>
<a href="/messeges/">Back to Messeges!</a>
</body>
</html>

{{define "UIkit"}}
<!-- UIkit CSS -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/css/uikit.min.css" />

<!-- UIkit JS -->
<script src="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/js/uikit.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/js/uikit-icons.min.js"></script>
{{end}}

{{define "JS"}}
<script>
    async function updateMessege(id){
        console.log('updateTask()');
        let taskForm = document.querySelector(`div[post-id="${id}"]`);
        let postName = taskForm.querySelector('input[name="Name"]').value;
        let postText = taskForm.querySelector('input[name="Text"]').value;

        let data = await fetch(`/messege/${id}`, {
                    method: 'PUT',
                    body: JSON.stringify({
                        name: postName,
                        text: postText,
                    }),
                });

        let dataTask = await data.json();
                if(dataTask){
                    console.log(dataTask);
                    window.location.reload();
                }
    }

    async function deleteMessege(id){
        console.log('deleteMessege()');

fetch(`/messege/${id}`, {
            method: 'DELETE',
        }).then(response => {
            window.location.replace('/messeges/');
        });
}

</script>
{{end}} 