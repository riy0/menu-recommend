import axios from 'axios'

class MenuApi{
    constructor(){
        this.apiBase = 'http://127.0.0.1:8080/menus'
    }

    menus(){
        return axios
        .get(`${this.apiBase}`)
        .then(json => {
            return json
        })
        .catch(e => ({ error: e}))
    }

    menu(id){
        return axios
            .get(`http://127.0.0.1:8080/menus/${id}`)
            .then(json => {
                return json
            })
                .catch({ error: e } ))
    }
}

const menuApi = new MenuApi()

export default menuApi
