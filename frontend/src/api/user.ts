export const HOST_URL = 'http://0.0.0.0:8000/api'

export const headers = {
  'Content-type': 'application/json; charset=UTF-8',
  'Access-Control-Allow-Origin': '*',
  'Access-Control-Allow-Methods': '*',
}
export const login = async (
  email: string,
  password: string
): Promise<string | null> => {
  try {
    const { status } = await fetch(HOST_URL + '/user/login', {
      method: 'post',
      body: JSON.stringify({
        email: email,
        password: password,
      }),
    })
    if (status == 200) {
      return 'Kabi'
    }
  } catch (error) {
    console.log(error)
  }

  return null
}

export const signup = async (
  email: string,
  password: string,
  username: string
): Promise<string | null> => {
  try {
    const { status } = await fetch(HOST_URL + '/user/signup', {
      method: 'post',
      body: JSON.stringify({
        email: email,
        password: password,
        username: username,
      }),
    })
    if (status == 200) {
      return login(email, password)
    }
  } catch (error) {
    console.log(error)
  }

  return null
}
