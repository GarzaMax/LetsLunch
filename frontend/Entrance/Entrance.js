import React, {useState, useEffect, useContext} from 'react'
import RegistrationButton from "../Button/RegistrationButton";
import './entranceStyle.css'
import Validation from "../Validation/Validation";
import axios from "axios";
import tokenContext from "../tokenContext";
import personContext from "../personContext";
import {useNavigate} from "react-router-dom";

function Entrance() {
    const { token, setToken} = useContext(tokenContext)
    const [error, setError] = useState('')
    const {person, setPerson} = useContext(personContext)
    const navigate = useNavigate()


    const {blurHandler,
        phoneNumberHandler,
        passwordHandler,
        phoneNumber,
        phoneNumberDirty,
        phoneNumberError,
        formValid,
        focusedPhoneNumber,
        focusedData,
        password,
        focusedPassword,
        passwordDirty,
        passwordError} = Validation(setError)

    async function fetchEntrance() {
        try {
            console.log(token)
            const tokens = await axios.post('http://localhost:3333/api/user/token',
                {
                    phoneNumber: phoneNumber,
                    password: password
                });

            const response = await axios.get('http://localhost:3333/api/user/find')
            setPerson(response.data.filter(elem => elem.phoneNumber === phoneNumber))
            console.log(response)

            axios.defaults.headers.common['Authorization'] = token;
        } catch (error) {
            console.log(error);
            if (error.response.status === 422) {
                setError('Такой номер телефона уже используется')
            }
        }
    }

    function onclick() {
        if (formValid) {
            fetchEntrance()
            navigate('/mainpage')
        }
    }

    return(
        <div className="entrance">
            <form style={{textAlign: 'center', display: 'flex-block'}}>
                <h1>Let's Lunch</h1>
                <h2>Войти</h2>

                {<div className="numberPhone" style={{color: '#2a1100', textAlign: 'left'}}>Номер телефона</div>}
                <input className="Input"
                       onChange={e => phoneNumberHandler(e)}
                       value={phoneNumber}
                       onBlur={e => blurHandler(e)}
                       focused={focusedData.toString()}
                       name='phone'
                       type="text"
                       placeholder='Enter your phone number...'
                       style={{border: (focusedPhoneNumber && phoneNumberError) ? '1px solid red' : '1px solid #ccc'}}/>
                {(phoneNumberDirty && phoneNumberError) && <div style={{color: 'red'}}>{phoneNumberError}</div>}


                {<div className="numberPhone" style={{color: '#2a1100', textAlign: 'left'}}>Пароль</div>}
                <input className="Input"
                       onChange={e => passwordHandler(e)}
                       value={password}
                       onBlur={e => blurHandler(e)}
                       focused={focusedPassword.toString()}
                       name='password'
                       type="password"
                       placeholder='Enter your password...'
                       style={{border: (focusedPassword && passwordError) ? '1px solid red' : '1px solid #ccc'}}/>
                {(passwordDirty && passwordError) && <div style={{color: 'red'}}>{passwordError}</div>}

                <div className="reg" style={{textAlign: 'center'}}>
                    <button className="in-website-from-entrance" onClick={e => onclick(e.preventDefault())}
                            disabled={!formValid}>Войти
                    </button>
                </div>


                    <RegistrationButton>Registration</RegistrationButton>
            </form>
        </div>
)
}

export default Entrance;