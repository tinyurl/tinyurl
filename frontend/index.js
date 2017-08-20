var apiUrl="http://localhost:8877"

app = new Vue({
    el: "#app",
    data: {
        url: ''
    },
    methods: {
        validateInput: function(url) {
            // refer http://www.cnblogs.com/554006164/archive/2009/06/16/1504160.html
            var urlRegex = "^((https|http)?://)"  
                + "(([0-9]{1,3}\.){3}[0-9]{1,3}" // IP形式的URL- 199.194.52.184  
                + "|" // 允许IP和DOMAIN（域名） 
                + "([0-9a-z_!~*'()-]+\.)*" // 域名- www.  
                + "([0-9a-z][0-9a-z-]{0,61})?[0-9a-z]\." // 二级域名  
                + "[a-z]{2,6})" // first level domain- .com or .museum  
                + "(:[0-9]{1,4})?" // 端口- :80  
                + "((/?)|" // a slash isn't required if there is no file name  
                + "(/[0-9a-z_!~*'().;?:@&=+$,%#-]+)+/?)$"
            var re = new RegExp(urlRegex)
            if (re.test(url)) {
                console.log('url ', url, ' is validate')
                return true
            }
            console.log('url ', url, ' is not validate')            
            return false
        },
        shortenUrl: function() {
            self = this
            if (self.url == "") {
                alert("please input correct url")
                return
            }

            if (!self.validateInput(self.url)) {
                alert("your input url is not validate!")
                return
            }

            var postUrl = apiUrl+"/api/v1/shorten"
            config = { headers: { 'Content-Type': 'multipart/form-data'}}
            formData = new FormData()
            formData.append('origin_url', self.url)
            
            axios.post(postUrl, formData, config)
            .then(function(response) {
                data = response.data
                self.url="http://tinyurl.adolphlwq.xyz/n/"+data["short_path"]
            })
            .catch(function(error) {
                alert(error)
            })
        }
    }
})
