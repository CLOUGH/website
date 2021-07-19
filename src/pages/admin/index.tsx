import React from 'react'
import { useAuth } from '../../context/AuthUserContext'

export default function AdminDashboard(){
  const {authUser} = useAuth()
  return (
    <div>
      <p>Your email is {authUser?.email}.</p>
    </div>
  )
}