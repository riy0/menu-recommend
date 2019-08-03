import axios from 'axios'

class MenuApi{
    constructor(){
        this.apiBase = process.env.baseUrl
    }

    menus(){
        console.log(process.env.baseUrl) 
        return axios
        .get(`${this.apiBase}`)
        .then(json => {
            return json
        })
        .catch(e => ({ error: e}))
    }

    menu(id){
        return axios
            .get(`${menuUrl}`)
            .then(json => {
                return json
            })
                .catch({ error: e } ))
    }
}

const menuApi = new MenuApi()

export default menuApi
