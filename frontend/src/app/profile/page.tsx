"use client"

import { useState } from "react"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import * as z from "zod"
import { Loader2, User as UserIcon } from "lucide-react"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { api } from "@/lib/axios"
import { useAuthStore } from "@/store/auth"

const profileSchema = z.object({
  name: z.string().min(2, { message: "Name must be at least 2 characters" }),
  avatar_url: z.string().url({ message: "Must be a valid URL" }).optional().or(z.literal("")),
})

const passwordSchema = z.object({
  currentPassword: z.string().min(1, { message: "Current password is required" }),
  newPassword: z.string().min(8, { message: "New password must be at least 8 characters" }),
  confirmNewPassword: z.string()
}).refine((data) => data.newPassword === data.confirmNewPassword, {
  message: "Passwords don't match",
  path: ["confirmNewPassword"],
})

export default function ProfilePage() {
  const { user, setUser } = useAuthStore()
  
  const [profileMsg, setProfileMsg] = useState({ type: "", text: "" })
  const [passwordMsg, setPasswordMsg] = useState({ type: "", text: "" })

  const profileForm = useForm<z.infer<typeof profileSchema>>({
    resolver: zodResolver(profileSchema),
    defaultValues: { 
      name: user?.name || "",
      avatar_url: user?.avatar_url || ""
    }
  })

  const passwordForm = useForm<z.infer<typeof passwordSchema>>({
    resolver: zodResolver(passwordSchema),
  })

  const onProfileSubmit = async (data: z.infer<typeof profileSchema>) => {
    setProfileMsg({ type: "", text: "" })
    try {
      const payload = {
        name: data.name,
        avatar_url: data.avatar_url || null
      }
      const response = await api.put("/profile", payload)
      setUser({ ...user!, name: response.data.data.name, avatar_url: response.data.data.avatar_url })
      setProfileMsg({ type: "success", text: "Profile updated successfully" })
    } catch (err: any) {
      setProfileMsg({ type: "error", text: err.response?.data?.message || "Failed to update profile" })
    }
  }

  const onPasswordSubmit = async (data: z.infer<typeof passwordSchema>) => {
    setPasswordMsg({ type: "", text: "" })
    try {
      await api.post("/auth/change-password", {
        currentPassword: data.currentPassword,
        newPassword: data.newPassword
      })
      setPasswordMsg({ type: "success", text: "Password changed successfully" })
      passwordForm.reset()
    } catch (err: any) {
      setPasswordMsg({ type: "error", text: err.response?.data?.message || "Failed to change password" })
    }
  }

  if (!user) return null

  return (
    <div className="container max-w-4xl mx-auto py-10 px-4 space-y-8">
      <div className="flex items-center space-x-4">
        <div className="bg-primary/10 p-4 rounded-full flex items-center justify-center h-16 w-16 overflow-hidden">
          {user.avatar_url ? (
            <img src={user.avatar_url} alt={user.name} className="h-full w-full object-cover rounded-full" />
          ) : (
            <UserIcon className="h-8 w-8 text-primary" />
          )}
        </div>
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Account Profile</h1>
          <p className="text-muted-foreground">Manage your account settings and password.</p>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <Card>
          <CardHeader>
            <CardTitle>Profile Details</CardTitle>
            <CardDescription>Update your personal information.</CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={profileForm.handleSubmit(onProfileSubmit)} className="space-y-4">
              {profileMsg.text && (
                <div className={`p-3 text-sm rounded-md ${profileMsg.type === "error" ? "bg-destructive/10 text-destructive" : "bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400"}`}>
                  {profileMsg.text}
                </div>
              )}
              <div className="space-y-2">
                <Label htmlFor="email">Email</Label>
                <Input id="email" value={user.email} disabled />
              </div>
              <div className="space-y-2">
                <Label htmlFor="name">Name</Label>
                <Input id="name" {...profileForm.register("name")} />
                {profileForm.formState.errors.name && (
                  <p className="text-sm text-destructive">{profileForm.formState.errors.name.message}</p>
                )}
              </div>
              <div className="space-y-2">
                <Label htmlFor="avatar_url">Avatar URL (Optional)</Label>
                <Input id="avatar_url" placeholder="https://example.com/avatar.png" {...profileForm.register("avatar_url")} />
                {profileForm.formState.errors.avatar_url && (
                  <p className="text-sm text-destructive">{profileForm.formState.errors.avatar_url.message}</p>
                )}
              </div>
              <Button type="submit" disabled={profileForm.formState.isSubmitting}>
                {profileForm.formState.isSubmitting && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                Save Changes
              </Button>
            </form>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Change Password</CardTitle>
            <CardDescription>Update your password to keep your account secure.</CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={passwordForm.handleSubmit(onPasswordSubmit)} className="space-y-4">
              {passwordMsg.text && (
                <div className={`p-3 text-sm rounded-md ${passwordMsg.type === "error" ? "bg-destructive/10 text-destructive" : "bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400"}`}>
                  {passwordMsg.text}
                </div>
              )}
              <div className="space-y-2">
                <Label htmlFor="currentPassword">Current Password</Label>
                <Input id="currentPassword" type="password" {...passwordForm.register("currentPassword")} />
                {passwordForm.formState.errors.currentPassword && (
                  <p className="text-sm text-destructive">{passwordForm.formState.errors.currentPassword.message}</p>
                )}
              </div>
              <div className="space-y-2">
                <Label htmlFor="newPassword">New Password</Label>
                <Input id="newPassword" type="password" {...passwordForm.register("newPassword")} />
                {passwordForm.formState.errors.newPassword && (
                  <p className="text-sm text-destructive">{passwordForm.formState.errors.newPassword.message}</p>
                )}
              </div>
              <div className="space-y-2">
                <Label htmlFor="confirmNewPassword">Confirm New Password</Label>
                <Input id="confirmNewPassword" type="password" {...passwordForm.register("confirmNewPassword")} />
                {passwordForm.formState.errors.confirmNewPassword && (
                  <p className="text-sm text-destructive">{passwordForm.formState.errors.confirmNewPassword.message}</p>
                )}
              </div>
              <Button type="submit" variant="secondary" disabled={passwordForm.formState.isSubmitting}>
                {passwordForm.formState.isSubmitting && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                Update Password
              </Button>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
