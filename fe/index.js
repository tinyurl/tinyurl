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
            var postUrl = apiUrl+"/api/v1/shorten?longurl="+this.url
            var self = this
            this.$http.post(postUrl).then(response=>{
                ret = response.body
                this.url="http://tinyurl.adolphlwq.xyz/"+ret["shortpath"]
            }, response=>{
                alert("shorten url error")
            })
        }
    }
})
