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
            var shortpath

            axios.post(postUrl, {
                "longurl": this.url
            })
            .then(function(response) {
                //this.url="http://tinyurl.adolphlwq.xyz/n/"+shortpath
                console.log(response)
                this.url='sb'
            })
            .catch(function(error) {
                alert(error)
            })
        }
    }
})
