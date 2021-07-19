import React, { useState } from "react";
import Image from 'next/image';
import styled from "styled-components";
import { useForm } from "react-hook-form";
import { useAuth } from "../context/AuthUserContext";
import { useRouter } from "next/dist/client/router";

const LoginPageStyledWrapper = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: center;
    min-height: 100vh;
    align-items: center;
    
    form {
       width: 300px;
       text-align: center;
        input[type="email"] {
            margin-bottom: -1px;
            border-bottom-right-radius: 0;
            border-bottom-left-radius: 0;
        }

        input[type="password"] {
            margin-bottom: 10px;
            border-top-left-radius: 0;
            border-top-right-radius: 0;
        }
        .form-floating:focus-within {
            z-index: 2;
        }
        .checkbox {
           font-weight: 400;
        }

    } 
`;

interface LoginCredentials {
    email: string;
    password: string;
}

export default function Login() {
    const {signInWithEmailAndPassword, authUser} = useAuth();
    const { register, handleSubmit, watch, formState: { errors } } = useForm();
    const [loginError, setLoginError] = useState<any>(null);
    const router = useRouter();


    const onSubmit = (loginCredentials: LoginCredentials) => {
        console.log({loginCredentials});
        setLoginError(null);
        signInWithEmailAndPassword(loginCredentials.email, loginCredentials.password)
            .then(authUser => {
                console.log({authUser})
                router.push('/admin');
            }).catch(error => {
                console.error({error})
                setLoginError(error);
            });
    };

    return (
        <LoginPageStyledWrapper className="bg-light">
            <form onSubmit={handleSubmit(onSubmit)}>
                <h1>WC</h1>
                <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
                {loginError && (
                    <div className="alert alert-danger">
                        {loginError.message}
                    </div>
                )}

                <div className="form-floating">
                    <input type="email"
                        className="form-control"
                        id="floatingInput"
                        placeholder="name@example.com"
                        {...register('email', { required: true })}
                    />
                    <label>Email address</label>
                </div>

                <div className="form-floating">
                    <input type="password"
                        className="form-control"
                        placeholder="Password"
                        {...register('password', { required: true })}
                    />
                    <label>Password</label>
                </div>

                <div className="checkbox mb-3">
                    <label>
                        <input type="checkbox" value="remember-me" /> Remember me
                    </label>
                </div>
                <button className="w-100 btn btn-lg btn-primary" type="submit" >Sign in</button>
                <p className="mt-5 mb-3 text-muted">&copy; 2017–2021</p>
            </form>
        </LoginPageStyledWrapper>
    );
}