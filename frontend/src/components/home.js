import React from 'react'
import { useNavigate } from 'react-router-dom'
import Header from "./header";
import { Button } from '@chakra-ui/react'

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
            <Header />
            <div className={'titleContainer'}>
                {loggedIn ? <div>Welcome, {username}!</div> : <div>Welcome!</div>}
            </div>
            <div>This is the home page.</div>
            <div className={'buttonContainer'}>
                <Button colorScheme="teal" size="md" onClick={onButtonClick} className={'loginButton'}>
                    {loggedIn ? 'Log out' : 'Log in'}
                </Button>
            </div>
            {!loggedIn && (
                <div className='registerContainer'>
                    <div className='registerText'>Don't have an account?</div>
                    <Button colorScheme='teal' size='md' onClick={() => navigate('/register')} className={'loginButton'}>
                        Register
                    </Button>
                </div>
            )}
        </div>
    )
}

export default Home