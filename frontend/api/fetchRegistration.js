import axios from "axios";

async function fetchRegistration(data, phoneNumber, male) {
    const response = await axios.post('http://localhost:3333/api/user/create', {
        username: "NoName",
        displayName: "NoName",
        birthday: {data},
        phoneNumber: {phoneNumber},
        gender: {male}})
    /*setVisit(response.data)*/
}

export default fetchRegistration;