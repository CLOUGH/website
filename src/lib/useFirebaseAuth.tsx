import { UserInfo } from 'firebase-functions/lib/providers/auth';
import React, {useState, useEffect } from 'react';
import firebase from './firebase';

export default function useFirebaseAuth() {
    const [authUser, setAuthUser] = useState<any>(null);
    const [loading, setLoading] = useState<boolean>(false);

    const authStateChanged = async(authState: any) => {
        setLoading(true);
        
        if(!authState) {
            setLoading(false);
            return;
        }


        setAuthUser(authState)
        setLoading(false);
    }

    const signInWithEmailAndPassword = async (email: string, password: string) => {
        return firebase.auth().signInWithEmailAndPassword(email, password);
    }

    const signOut = async () => {
        firebase.auth().signOut().then(() => {
            setAuthUser(null);
            setLoading(true);
        });
    }

    useEffect(() => {
        const unsubscribe = firebase.auth().onAuthStateChanged(authStateChanged);
        return () => unsubscribe();
    }, []);


    return {
        authUser,
        loading,
        signInWithEmailAndPassword,
        signOut
    };
}