//var apiUrl="http://tinyurl.api.adolphlwq.xyz"
var apiUrl="http://localhost:8877"

app = new Vue({
    el: "#app",
    data: {
        url: ''
    },
    methods: {
        shortenUrl: function() {
            if (this.url == "") {
                alert("please input correct url")
                return
            }
            var postUrl = apiUrl+"/api/v1/shorten"
            formData = new FormData()
            formData.append('longurl', this.url)
            this.$http.post(postUrl, formData).then(response=>{
                ret = response.body
                this.url="http://tinyurl.adolphlwq.xyz/n/"+ret["shortpath"]
            }, response=>{
                alert("shorten url error")
            })
        }
    }
})
