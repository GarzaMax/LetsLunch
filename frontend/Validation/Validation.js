import React, {useEffect, useState} from 'react'

function Validation(setError) {
    const [phoneNumber, setphoneNumber] = useState('')
    const [data, setData] = useState('')
    const [name, setName] = useState('')
    const [password, setPassword] = useState('')
    const [phoneNumberDirty, setphoneNumberDirty] = useState(false)
    const [dataDirty, setDataDirty] = useState(false)
    const [nameDirty, setNameDirty] = useState(false)
    const [passwordDirty, setPasswordDirty] = useState(false)
    const [phoneNumberError, setphoneNumberError] = useState('Номер не может быть пустым')
    const [dataError, setDataError] = useState('Дата не может быть пустой')
    const [nameError, setNameError] = useState('Имя не может быть пустым')
    const [passwordError, setPasswordError] = useState('Пароль не может быть пустым')
    const [formValid, setFormValid] = useState(false)
    const [focusedPhoneNumber, setFocusedPhoneNumber] = useState(false)
    const [focusedData, setFocusedData] = useState(false)
    const [focusedName, setFocusedName] = useState(false)
    const [focusedPassword, setFocusedPassword] = useState(false)

    const phoneNumberHandler = (e) => {
        setphoneNumber(e.target.value)
        const re = /^[\\+]?[(]?[0-9]{3}[)]?[-\s\\.]?[0-9]{3}[-\s\\.]?[0-9]{4,6}$/i;
        if (!re.test(String(e.target.value).toLowerCase())) {
            setphoneNumberError('Некорректный номер')
            setError('')
        } else {
            setphoneNumberError('')
            setError('')
        }
    }

    const dataHandler = (e) => {
        setData(e.target.value)
        if (e.target.value.length !== 0) {
            setDataError('')
        }
    }

    const nameHandler = (e) => {
        setName(e.target.value)
        if (e.target.value.length !== 0) {
            setNameError('')
        }
    }

    const passwordHandler = (e) => {
        setPassword(e.target.value)
        if (e.target.value.length !== 0) {
            setPasswordError('')
        }
    }

    useEffect(() => {
        if (phoneNumberError || passwordError) {
            setFormValid(false)
        } else {
            setFormValid(true)
        }
    }, [phoneNumberError, passwordError])


    const blurHandler = (e) => {
        switch (e.target.name) {
            case 'phone':
                setphoneNumberDirty(true)
                setFocusedPhoneNumber(true)
                break
            case 'date':
                setDataDirty(true)
                setFocusedData(true)
                break
            case 'name':
                setNameDirty(true)
                setFocusedName(true)
                break
            case 'password':
                setPasswordDirty(true)
                setFocusedPassword(true)
                break
            default:
                break
        }
    }

    return {blurHandler,
        phoneNumberHandler,
        nameHandler,
        passwordHandler,
        dataHandler,
        name,
        phoneNumber,
        setphoneNumber,
        phoneNumberDirty,
        setphoneNumberDirty,
        phoneNumberError,
        setphoneNumberError,
        formValid,
        setFormValid,
        focusedPhoneNumber,
        setFocusedPhoneNumber,
        focusedData,
        dataDirty,
        dataError,
        data,
        nameDirty,
        focusedName,
        nameError,
        password,
        focusedPassword,
        passwordDirty,
        passwordError,
    };
}

export default Validation;