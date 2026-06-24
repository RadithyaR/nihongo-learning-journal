"use client"

import { useState } from "react"
import { api } from "@/lib/axios"
import { 
  BookOpen, 
  Eye, 
  RefreshCw, 
  ArrowRight,
  BrainCircuit,
  Type,
  FileText
} from "lucide-react"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"

type ReviewMode = "vocabulary" | "kanji" | "grammar" | null

interface ReviewItem {
  id: string
  question: string
  reading?: string
  meaning?: string
  extra?: string
}

export default function ReviewPage() {
  const [mode, setMode] = useState<ReviewMode>(null)
  const [currentItem, setCurrentItem] = useState<ReviewItem | null>(null)
  const [showAnswer, setShowAnswer] = useState(false)
  const [loading, setLoading] = useState(false)
  const [sessionCount, setSessionCount] = useState(0)
  const [errorMsg, setErrorMsg] = useState<string | null>(null)

  const fetchNextItem = async (selectedMode: ReviewMode) => {
    if (!selectedMode) return
    setLoading(true)
    setShowAnswer(false)
    try {
      let endpoint = ""
      if (selectedMode === "vocabulary") endpoint = "/reviews/next"
      if (selectedMode === "kanji") endpoint = "/reviews/kanji/next"
      if (selectedMode === "grammar") endpoint = "/reviews/grammar/next"

      const response = await api.get(endpoint)
      const data = response.data.data

      if (selectedMode === "vocabulary") {
        setCurrentItem({ id: data.id, question: data.word, reading: data.reading })
      } else if (selectedMode === "kanji") {
        setCurrentItem({ 
          id: data.id, 
          question: data.character, 
          meaning: data.meaning, 
          reading: `${data.onyomi ? `On: ${data.onyomi}` : ''} ${data.kunyomi ? `Kun: ${data.kunyomi}` : ''}`
        })
      } else if (selectedMode === "grammar") {
        setCurrentItem({ id: data.id, question: data.pattern, meaning: data.meaning })
      }
    } catch (err: any) {
      if (err.response?.status === 404) {
        // No more items
        setCurrentItem(null)
        if (err.response.data?.message?.includes("You haven't added any")) {
          setErrorMsg(err.response.data.message)
        } else {
          setErrorMsg(null)
        }
      }
      console.error("Failed to load next review", err)
    } finally {
      setLoading(false)
    }
  }

  const startSession = (selectedMode: ReviewMode) => {
    setMode(selectedMode)
    setSessionCount(0)
    setErrorMsg(null)
    fetchNextItem(selectedMode)
  }

  const submitReview = async (rating: "AGAIN" | "HARD" | "GOOD" | "EASY") => {
    if (!currentItem || !mode) return
    setLoading(true)

    try {
      let endpoint = ""
      let payload = {}
      
      if (mode === "vocabulary") {
        endpoint = "/reviews"
        payload = { item_id: currentItem.id, rating }
      } else if (mode === "kanji") {
        endpoint = "/reviews/kanji"
        payload = { item_id: currentItem.id, rating }
      } else if (mode === "grammar") {
        endpoint = "/reviews/grammar"
        payload = { grammar_id: currentItem.id, rating }
      }

      await api.post(endpoint, payload)
      setSessionCount(prev => prev + 1)
      fetchNextItem(mode)
    } catch (err) {
      console.error("Failed to submit review", err)
      setLoading(false)
    }
  }

  if (!mode) {
    return (
      <div className="container max-w-4xl mx-auto py-12 px-4">
        <div className="text-center mb-12">
          <h1 className="text-4xl font-bold tracking-tight mb-4">Study Session</h1>
          <p className="text-xl text-muted-foreground">What would you like to review today?</p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <Card className="hover:border-primary/50 transition-colors cursor-pointer" onClick={() => startSession("vocabulary")}>
            <CardHeader className="text-center pb-2">
              <Type className="h-12 w-12 mx-auto text-primary mb-4" />
              <CardTitle>Vocabulary</CardTitle>
            </CardHeader>
            <CardContent className="text-center text-muted-foreground text-sm">
              Review words, readings, and meanings using Spaced Repetition.
            </CardContent>
          </Card>

          <Card className="hover:border-primary/50 transition-colors cursor-pointer" onClick={() => startSession("kanji")}>
            <CardHeader className="text-center pb-2">
              <BookOpen className="h-12 w-12 mx-auto text-primary mb-4" />
              <CardTitle>Kanji</CardTitle>
            </CardHeader>
            <CardContent className="text-center text-muted-foreground text-sm">
              Practice characters, Onyomi, and Kunyomi readings.
            </CardContent>
          </Card>

          <Card className="hover:border-primary/50 transition-colors cursor-pointer" onClick={() => startSession("grammar")}>
            <CardHeader className="text-center pb-2">
              <BrainCircuit className="h-12 w-12 mx-auto text-primary mb-4" />
              <CardTitle>Grammar</CardTitle>
            </CardHeader>
            <CardContent className="text-center text-muted-foreground text-sm">
              Master sentence patterns and Japanese grammar rules.
            </CardContent>
          </Card>
        </div>
      </div>
    )
  }

  return (
    <div className="container max-w-3xl mx-auto py-10 px-4 flex flex-col items-center min-h-[80vh]">
      <div className="w-full flex justify-between items-center mb-8">
        <Button variant="ghost" onClick={() => setMode(null)}>
          &larr; Back to Menu
        </Button>
        <div className="text-sm font-medium bg-muted px-3 py-1 rounded-full">
          Session count: {sessionCount}
        </div>
      </div>

      {loading && !currentItem ? (
        <div className="flex-1 flex flex-col items-center justify-center space-y-4">
          <RefreshCw className="h-12 w-12 animate-spin text-primary/50" />
          <p className="text-muted-foreground">Loading next item...</p>
        </div>
      ) : !currentItem ? (
        <Card className={`w-full text-center py-12 ${errorMsg ? 'bg-red-500/10 border-red-500/20' : 'bg-green-500/10 border-green-500/20'}`}>
          <CardContent className="space-y-4">
            <div className={`h-16 w-16 rounded-full flex items-center justify-center mx-auto mb-4 ${errorMsg ? 'bg-red-500/20' : 'bg-green-500/20'}`}>
              <FileText className={`h-8 w-8 ${errorMsg ? 'text-red-600 dark:text-red-400' : 'text-green-600 dark:text-green-400'}`} />
            </div>
            <h2 className={`text-2xl font-bold ${errorMsg ? 'text-red-700 dark:text-red-400' : 'text-green-700 dark:text-green-400'}`}>
              {errorMsg ? "No Data Available" : "All Caught Up!"}
            </h2>
            <p className="text-muted-foreground max-w-md mx-auto">
              {errorMsg ? errorMsg : `You have no more pending ${mode} reviews at the moment. Great job keeping up with your studies!`}
            </p>
            <Button onClick={() => setMode(null)} className="mt-4" variant="outline">
              Return to Menu
            </Button>
          </CardContent>
        </Card>
      ) : (
        <Card className="w-full shadow-lg border-primary/10">
          <CardHeader className="text-center pb-8 pt-12 border-b">
            <CardDescription className="uppercase tracking-widest font-semibold text-primary mb-2">
              {mode}
            </CardDescription>
            <CardTitle className="text-6xl sm:text-7xl font-bold">
              {currentItem.question}
            </CardTitle>
          </CardHeader>
          
          <CardContent className="pt-8 pb-12 min-h-[200px] flex flex-col items-center justify-center text-center">
            {!showAnswer ? (
              <Button 
                onClick={() => setShowAnswer(true)} 
                size="lg" 
                className="text-lg px-8 h-14 w-full sm:w-auto"
              >
                <Eye className="mr-2 h-5 w-5" />
                Show Answer
              </Button>
            ) : (
              <div className="space-y-6 animate-in fade-in zoom-in duration-300 w-full">
                {currentItem.reading && (
                  <div>
                    <h3 className="text-sm uppercase tracking-wider text-muted-foreground mb-1">Reading</h3>
                    <p className="text-2xl font-medium">{currentItem.reading}</p>
                  </div>
                )}
                {currentItem.meaning && (
                  <div>
                    <h3 className="text-sm uppercase tracking-wider text-muted-foreground mb-1">Meaning</h3>
                    <p className="text-2xl font-medium">{currentItem.meaning}</p>
                  </div>
                )}
              </div>
            )}
          </CardContent>

          {showAnswer && (
            <CardFooter className="flex flex-wrap sm:flex-nowrap justify-center gap-2 sm:gap-4 pt-6 border-t bg-muted/30">
              <Button 
                variant="destructive" 
                className="w-full sm:flex-1 h-12 text-sm sm:text-base font-semibold"
                onClick={() => submitReview("AGAIN")}
                disabled={loading}
              >
                Again (1m)
              </Button>
              <Button 
                variant="outline" 
                className="w-full sm:flex-1 h-12 text-sm sm:text-base font-semibold border-orange-200 text-orange-600 hover:bg-orange-50 hover:text-orange-700 dark:border-orange-900/50 dark:text-orange-400 dark:hover:bg-orange-900/20"
                onClick={() => submitReview("HARD")}
                disabled={loading}
              >
                Hard
              </Button>
              <Button 
                variant="outline" 
                className="w-full sm:flex-1 h-12 text-sm sm:text-base font-semibold border-blue-200 text-blue-600 hover:bg-blue-50 hover:text-blue-700 dark:border-blue-900/50 dark:text-blue-400 dark:hover:bg-blue-900/20"
                onClick={() => submitReview("GOOD")}
                disabled={loading}
              >
                Good
              </Button>
              <Button 
                variant="outline" 
                className="w-full sm:flex-1 h-12 text-sm sm:text-base font-semibold border-green-200 text-green-600 hover:bg-green-50 hover:text-green-700 dark:border-green-900/50 dark:text-green-400 dark:hover:bg-green-900/20"
                onClick={() => submitReview("EASY")}
                disabled={loading}
              >
                Easy
              </Button>
            </CardFooter>
          )}
        </Card>
      )}
    </div>
  )
}
