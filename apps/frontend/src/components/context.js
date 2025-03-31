import { createContext } from 'react'

const UserObject = {
    users: [],
    page: 1,
    Length: 10,
    PageSize: 20
}

export default createContext(UserObject)