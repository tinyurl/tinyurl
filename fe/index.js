//var apiUrl="http://tinyurl.api.adolphlwq.xyz"
var apiUrl="http://localhost:8877"

app = new Vue({
    el: "#app",
    data: {
        url: ''
    },
    methods: {
        shortenUrl: function() {
            self = this
            if (self.url == "") {
                alert("please input correct url")
                return
            }

            var postUrl = apiUrl+"/api/v1/shorten"
            config = { headers: { 'Content-Type': 'multipart/form-data'}}
            formData = new FormData()
            formData.append('longurl', self.url)
            
            axios.post(postUrl, formData, config)
            .then(function(response) {
                data = response.data
                self.url="http://tinyurl.adolphlwq.xyz/n/"+data["shortpath"]
            })
            .catch(function(error) {
                alert(error)
            })
        }
    }
})
