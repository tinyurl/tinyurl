//var apiUrl="http://tinyurl.api.adolphlwq.xyz"
var apiUrl="http://localhost:8877"

app = new Vue({
    el: "#app",
    data: {
        url: ''
    },
    methods: {
        validateInput: function() {
            // refer http://www.cnblogs.com/554006164/archive/2009/06/16/1504160.html
            var urlRegex = "^((https|http|ftp|rtsp|mms)?://)"  
                + "?(([0-9a-z_!~*'().&=+$%-]+: )?[0-9a-z_!~*'().&=+$%-]+@)?" //ftp的user@  
                + "(([0-9]{1,3}\.){3}[0-9]{1,3}" // IP形式的URL- 199.194.52.184  
                + "|" // 允许IP和DOMAIN（域名） 
                + "([0-9a-z_!~*'()-]+\.)*" // 域名- www.  
                + "([0-9a-z][0-9a-z-]{0,61})?[0-9a-z]\." // 二级域名  
                + "[a-z]{2,6})" // first level domain- .com or .museum  
                + "(:[0-9]{1,4})?" // 端口- :80  
                + "((/?)|" // a slash isn't required if there is no file name  
                + "(/[0-9a-z_!~*'().;?:@&=+$,%#-]+)+/?)$"
            var re = new RegExp(urlRegex)
            if (re.test(this.url)) {
                return true
            }
            return false
        },
        shortenUrl: function() {
            self = this
            if (self.url == "") {
                alert("please input correct url")
                return
            }

            if (!self.validateInput()) {
                alert("your input url is not validate!")
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
