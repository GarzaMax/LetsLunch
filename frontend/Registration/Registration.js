import React, {useState} from 'react'
import EntranceButton from "../Button/EntranceButton";
import Validation from "../Validation/Validation";
import './Registration.css'
import axios from "axios";
import {useNavigate} from "react-router-dom";

function Registration() {
    const [male, setMale] = useState('')
    const [error, setError] = useState('')
    const navigate = useNavigate()

    const {blurHandler,
        phoneNumberHandler,
        dataHandler,
        nameHandler,
        passwordHandler,
        phoneNumber,
        phoneNumberDirty,
        phoneNumberError,
        formValid,
        focusedPhoneNumber,
        focusedData,
        dataDirty,
        dataError,
        data,
        name,
        nameDirty,
        focusedName,
        nameError,
        password,
        passwordError,
        passwordDirty,
        focusedPassword} = Validation(setError)

    async function fetchRegistration() {
        let d1 = new Date(data.getFullYear, data.getMonth, data.getDay);
        try {
            const response = await axios.post('http://localhost:3333/api/user/create',
                {
                    username: name,
                    displayName: "NoName",
                    birthday: d1,
                    phoneNumber: phoneNumber,
                    password: password,
                    gender: male,
                });
        } catch (error) {
            console.log(error);
            if (error.response.status === 422) {
                setError('Такой номер телефона уже используется')
            }
        }
    }

    function onclick() {
        if (formValid) {
            fetchRegistration()
        }
        if (formValid && !error) {
            navigate('/')
        }
    }

    return (
        <div className="registration">
            <form style={{textAlign: 'center', display: 'flex-block'}}>
                <h1>Let's Lunch</h1>
                <h2>Регистрация</h2>


                {<div className="numberPhone" style={{color: '#2a1100', textAlign: 'left'}}>Имя</div>}
                <input className="Input"
                       onChange={e => nameHandler(e)}
                       value={name}
                       onBlur={e => blurHandler(e)}
                       focused={focusedData.toString()}
                       name='name'
                       type="text"
                       placeholder='Enter your name...'
                style={{border: (focusedName && nameError) ? '1px solid red' : '1px solid #ccc'}}/>
                {(nameDirty && nameError) && <div style={{color: 'red'}}>{nameError}</div>}


                {<div className="numberPhone" style={{color: '#2a1100', textAlign: 'left'}}>Номер телефона</div>}
                <input className="Input"
                       onChange={e => phoneNumberHandler(e)}
                       value={phoneNumber}
                       onBlur={e => blurHandler(e)}
                       focused={focusedPhoneNumber.toString()}
                       name='phone'
                       type="text"
                       placeholder='Enter your phone number...'
                       style={{border: ((focusedPhoneNumber && phoneNumberError) || error) ? '1px solid red' : '1px solid #ccc'}}/>
                {(phoneNumberDirty && phoneNumberError) && <div style={{color: 'red'}}>{phoneNumberError}</div>}
                {(error) && <div style={{color: 'red'}}>{error}</div>}

                {<div className="numberPhone" style={{textAlign: 'left'}}>Дата Рождения</div>}
                <input className="Input"
                       onChange={e => dataHandler(e)}
                       value={data}
                       onBlur={e => blurHandler(e)}
                       type="date"
                       focused={focusedData.toString()}
                       style={{border: (focusedData && dataError) ? '1px solid red' : '1px solid #ccc'}}/>
                {(dataDirty && dataError) && <div style={{color: 'red'}}>{dataError}</div>}


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


                {<div className="numberPhone" style={{color: '#2a1100', textAlign: 'left'}}>Пол</div>}
                <label style={{ color: '#2a1100',  fontSize: '1.5rem', display: "inline-block"}}><input  onClick={() => setMale(0)} style={{ color: '#2a1100',  fontSize: '1.5rem', display: "inline-block", marginLeft: '10px', marginRight: '5px'}} type="radio" name="gender" value="male"/>Мужской </label>
                <label style={{ color: '#2a1100',   fontSize: '1.5rem',  display: "inline-block"}}><input onClick={() => setMale(1)} style={{ color: '#2a1100',  fontSize: '1.5rem', display: "inline-block",  marginLeft: '10px',  marginRight: '5px'}} type="radio" name="gender" value="female"/>Женский</label>

                <div className="reg" style={{textAlign: 'center'}}>
                <button className="in-website" onClick={e => onclick(e.preventDefault())} disabled={!formValid}>Зарегистрироваться</button>
                <EntranceButton className="login-button">Войти</EntranceButton>
                </div>
            </form>
        </div>
    );
}

export default Registration;