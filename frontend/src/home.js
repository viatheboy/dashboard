import React from 'react'
import { useNavigate } from 'react-router-dom'

const Home = (props) => {
    const { loggedIn, username,  } = props
    const navigate = useNavigate()

    const onButtonClick = () => {
        if (loggedIn) {
            props.setLoggedIn(false)
            props.setUsername("")
        } else {
            navigate('/login')
        }
    }

    return (
        <div className="mainContainer">
            <div className={'titleContainer'}>
                {loggedIn ? <div>Welcome, {username}!</div> : <div>Welcome!</div>}
            </div>
            <div>This is the home page.</div>
            <div className={'buttonContainer'}>
                <input
                    className={'inputButton'}
                    type="button"
                    onClick={onButtonClick}
                    value={loggedIn ? 'Log out' : 'Log in'}
                />
            </div>
        </div>
    )
}

export default Home