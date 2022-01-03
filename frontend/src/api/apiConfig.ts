import axios from 'axios'

export const HOST_URL = 'http://0.0.0.0:8000/api'

export default axios.create({
  baseURL: HOST_URL,
  headers: {
    'Content-type': 'application/json',
    'Access-Control-Allow-Origin': 'true',
    'Access-Control-Allow-Methods': 'GET, PUT, POST, DELETE, OPTIONS',
  },
})
