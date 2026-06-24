"use client"

import { useEffect, useState } from "react"
import { api } from "@/lib/axios"
import { 
  BookOpen, 
  Flame, 
  GraduationCap, 
  Target, 
  AlertCircle,
  Clock,
  CheckCircle2,
  ListTodo
} from "lucide-react"

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import Link from "next/link"

interface DashboardData {
  totalVocabulary: number
  totalKanji: number
  totalGrammar: number
  reviewCountToday: number
  studyStreak: number
  activeGoals: number
  completedGoals: number
  dueToday: number
  overdue: number
  recentSessions: any[]
}

export default function DashboardPage() {
  const [data, setData] = useState<DashboardData | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchDashboard = async () => {
      try {
        const response = await api.get("/dashboard")
        setData(response.data.data)
      } catch (err) {
        console.error("Failed to load dashboard data", err)
      } finally {
        setLoading(false)
      }
    }

    fetchDashboard()
  }, [])

  if (loading) {
    return (
      <div className="flex flex-1 items-center justify-center min-h-[50vh]">
        <div className="animate-pulse space-y-4 text-center">
          <div className="h-12 w-12 bg-primary/20 rounded-full mx-auto" />
          <p className="text-muted-foreground">Loading dashboard...</p>
        </div>
      </div>
    )
  }

  return (
    <div className="container max-w-6xl mx-auto py-8 px-4 space-y-8">
      <div>
        <h1 className="text-3xl font-bold tracking-tight mb-2">Welcome Back!</h1>
        <p className="text-muted-foreground">Here is your learning progress at a glance.</p>
      </div>

      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <Card className="bg-primary/5 border-primary/20">
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Due Reviews</CardTitle>
            <ListTodo className="h-4 w-4 text-primary" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{data?.dueToday || 0}</div>
            <p className="text-xs text-muted-foreground mt-1">
              {data?.overdue ? (
                <span className="text-destructive font-semibold">{data.overdue} overdue items</span>
              ) : (
                "You're all caught up!"
              )}
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Study Streak</CardTitle>
            <Flame className="h-4 w-4 text-orange-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{data?.studyStreak || 0} Days</div>
            <p className="text-xs text-muted-foreground mt-1">
              {data?.reviewCountToday} items reviewed today
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Active Goals</CardTitle>
            <Target className="h-4 w-4 text-blue-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{data?.activeGoals || 0}</div>
            <p className="text-xs text-muted-foreground mt-1">
              {data?.completedGoals || 0} completed
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Items</CardTitle>
            <BookOpen className="h-4 w-4 text-green-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {(data?.totalVocabulary || 0) + (data?.totalKanji || 0) + (data?.totalGrammar || 0)}
            </div>
            <p className="text-xs text-muted-foreground mt-1">
              Vocab: {data?.totalVocabulary || 0} | Kanji: {data?.totalKanji || 0}
            </p>
          </CardContent>
        </Card>
      </div>

      <div className="grid gap-6 md:grid-cols-2">
        <Card className="col-span-1">
          <CardHeader>
            <CardTitle>Ready to Study?</CardTitle>
            <CardDescription>Start your daily review session now to keep up with your SRS schedule.</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="flex flex-col gap-3">
              <Link href="/review" className="w-full">
                <Button className="w-full h-12 text-lg">
                  <GraduationCap className="mr-2 h-5 w-5" />
                  Start Review Session
                </Button>
              </Link>
            </div>
            {((data?.dueToday || 0) > 0 || (data?.overdue || 0) > 0) && (
              <div className="flex items-center gap-2 p-3 bg-amber-500/10 text-amber-600 dark:text-amber-400 rounded-md text-sm">
                <AlertCircle className="h-4 w-4" />
                <span>You have pending reviews to complete today!</span>
              </div>
            )}
            {((data?.dueToday || 0) === 0 && (data?.overdue || 0) === 0) && (
              <div className="flex items-center gap-2 p-3 bg-green-500/10 text-green-600 dark:text-green-400 rounded-md text-sm">
                <CheckCircle2 className="h-4 w-4" />
                <span>Great job! You have completed all reviews for now.</span>
              </div>
            )}
          </CardContent>
        </Card>

        <Card className="col-span-1">
          <CardHeader>
            <CardTitle>Recent Sessions</CardTitle>
            <CardDescription>Your latest study activities.</CardDescription>
          </CardHeader>
          <CardContent>
            {data?.recentSessions && data.recentSessions.length > 0 ? (
              <div className="space-y-4">
                {data.recentSessions.slice(0, 4).map((session, i) => (
                  <div key={i} className="flex items-center gap-4 p-3 rounded-lg border bg-card/50">
                    <div className="bg-primary/10 p-2 rounded-full">
                      <Clock className="h-4 w-4 text-primary" />
                    </div>
                    <div className="flex-1 overflow-hidden">
                      <p className="text-sm font-medium truncate">{session.notes || "Study Session"}</p>
                      <p className="text-xs text-muted-foreground">
                        {new Date(session.sessionDate).toLocaleDateString()}
                      </p>
                    </div>
                  </div>
                ))}
              </div>
            ) : (
              <div className="text-center py-8 text-muted-foreground border-2 border-dashed rounded-lg">
                <p className="mb-2">No recent sessions found.</p>
                <Link href="/review">
                  <Button variant="link">Log your first session</Button>
                </Link>
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
