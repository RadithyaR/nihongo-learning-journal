"use client"

import { useEffect, useState } from "react"
import { useRouter, usePathname } from "next/navigation"
import { useAuthStore } from "@/store/auth"
import { Loader2 } from "lucide-react"

export function AuthGuard({ children }: { children: React.ReactNode }) {
  const router = useRouter()
  const pathname = usePathname()
  const { isAuthenticated } = useAuthStore()
  const [isMounted, setIsMounted] = useState(false)

  const publicRoutes = ["/login", "/register", "/forgot-password", "/reset-password"]
  const isPublicRoute = publicRoutes.includes(pathname)

  useEffect(() => {
    const initAuth = async () => {
      // If we don't have an access token, try to refresh silently using the httpOnly cookie
      if (!useAuthStore.getState().accessToken) {
        try {
          // The interceptor or a direct call to /auth/refresh
          const { api } = await import('@/lib/axios')
          const response = await api.post('/auth/refresh')
          const { access_token, user } = response.data.data
          useAuthStore.getState().setAuth(user, access_token, "")
        } catch (error) {
          // Silent failure, just means no valid session
        }
      }
      setIsMounted(true)
    }
    
    initAuth()
  }, [])

  useEffect(() => {
    if (!isMounted) return

    if (!isAuthenticated && !isPublicRoute && pathname !== "/") {
      router.push("/login")
    }

    if (isAuthenticated && isPublicRoute) {
      router.push("/dashboard")
    }
  }, [isAuthenticated, isPublicRoute, pathname, router, isMounted])

  if (!isMounted) {
    return (
      <div className="flex flex-1 items-center justify-center min-h-screen">
        <Loader2 className="h-8 w-8 animate-spin text-primary" />
      </div>
    )
  }

  // Prevent flash of protected content
  if (!isAuthenticated && !isPublicRoute && pathname !== "/") {
    return null
  }

  return <>{children}</>
}
