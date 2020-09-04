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
    {{.Title}} 
    <hr>
{{range .Messeges}}
<div>   
    <a href="/messege/{{.Id}}">{{.Text}}</a>
</div>
{{end}}
<div post-id="NewPost">
<input class="uk-input" type="text" name="Name" placeholder="Messege Name"/>
<input class="uk-input" type="text" name="Text" placeholder="Messege text"/>
<button class="uk-button" uk-button-default onclick="createTask()">Сохранить</button>
</div>
</body>

{{define "UIkit"}}
<!-- UIkit CSS -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/css/uikit.min.css" />

<!-- UIkit JS -->
<script src="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/js/uikit.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/js/uikit-icons.min.js"></script>
{{end}}

{{define "JS"}}
<script>
    async function createTask(){
        console.log('createTask()');
        let taskForm = document.querySelector('div[post-id="NewPost"]');
        let postName = taskForm.querySelector('input[name="Name"]').value;
        let postText = taskForm.querySelector('input[name="Text"]').value;
        console.log(postName, postText);
        console.log(JSON.stringify({name: postName,text: postText}));

        let data = await fetch('/messeges', {
                    method: 'POST',
                    body: JSON.stringify({
                        name: postName,
                        text: postText                        
                    }),
                });
        console.log(data);

                let dataTask = await data.json();
                if(dataTask){
                    console.log(dataTask);
                    window.location.reload();
                }
    }
</script>
{{end}} 