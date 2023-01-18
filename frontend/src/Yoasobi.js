import React,{useEffect,useState}  from 'react'
import axios from 'axios'
import './App.css'

function Yoasobi({id}){

    const [listSong,setListSong] = useState(null)
    
    const baseURL = `http://localhost:8080/get/${id}`
    
    const [email, setEmail] = useState(null)
    const [pass, setPass] = useState(null)


    useEffect(() =>{
        axios.get(baseURL) 
        .then((Response) => setListSong(Response.data))
        .catch(error => console.error(error));
    },[id]);

    return(
        <div>
            <div className ="BGIMG"></div>
            <div className ="table-menu">
                <div className ="BGcontainer">
                    <div className ="container">
                        <h2>Login</h2>
                        <form action="/Login" method="POST">

                            <input 
                                name="user_email" 
                                id="_email" 
                                className ="textbox" 
                                type="text" 
                                placeholder="Email" 
                                onChange={(e) => setEmail(e.target.value)}
                                rnodeequired
                            />

                            <input 
                                name="user_pass" 
                                id="_pass" 
                                className ="textbox" 
                                type="password" 
                                placeholder="Password" 
                                onChange={(e) => setPass(e.target.value)}
                                required
                            />
                            <button type="submit" className ="btn">Login</button>
                        </form>
                        <a href="/register.ejs">register</a>
                        <p>Check_Email: {email}</p>
                        <p>Check_Password: {pass}</p>
                    </div>
                </div>
            </div>
            {listSong ? 
                <a href={listSong.link}>
                    <p>{listSong.link}</p>
                </a> 
                : 
                <p>Loading...</p>
            }
        </div>
    )
}

export default Yoasobi