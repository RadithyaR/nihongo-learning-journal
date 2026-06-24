"use client"

import { useEffect, useState } from "react"
import { api } from "@/lib/axios"
import { Plus, Edit, Trash2, Loader2, Target, Calendar, CheckCircle2, Clock, XCircle } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Label } from "@/components/ui/label"

interface Goal {
  id: string
  title: string
  description?: string
  goalType?: string
  targetLevel?: string
  targetCount?: number
  currentCount: number
  progressPercentage: number
  targetDate: string
  status: string
}

export default function GoalsPage() {
  const [goals, setGoals] = useState<Goal[]>([])
  const [loading, setLoading] = useState(true)

  const [isDialogOpen, setIsDialogOpen] = useState(false)
  const [editId, setEditId] = useState<string | null>(null)
  
  const [formData, setFormData] = useState({
    title: "",
    description: "",
    goalType: "CUSTOM",
    targetLevel: "",
    targetCount: "",
    targetDate: ""
  })
  
  const [isSubmitting, setIsSubmitting] = useState(false)

  const fetchGoals = async () => {
    setLoading(true)
    try {
      const response = await api.get('/goals')
      // Ensure we get an array
      setGoals(Array.isArray(response.data.data) ? response.data.data : [])
    } catch (err) {
      console.error("Failed to load goals", err)
      setGoals([])
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchGoals()
  }, [])

  const deleteGoal = async (id: string) => {
    if (!confirm("Are you sure you want to delete this goal?")) return
    
    try {
      await api.delete(`/goals/${id}`)
      setGoals(goals.filter(g => g.id !== id))
    } catch (err) {
      console.error("Failed to delete", err)
    }
  }

  const openEdit = (goal: Goal) => {
    setFormData({
      title: goal.title,
      description: goal.description || "",
      goalType: goal.goalType || "CUSTOM",
      targetLevel: goal.targetLevel || "",
      targetCount: goal.targetCount ? goal.targetCount.toString() : "",
      // Format targetDate to YYYY-MM-DD for input type="date"
      targetDate: goal.targetDate ? new Date(goal.targetDate).toISOString().split('T')[0] : ""
    })
    setEditId(goal.id)
    setIsDialogOpen(true)
  }

  const handleOpenChange = (open: boolean) => {
    setIsDialogOpen(open)
    if (!open) {
      setEditId(null)
      setFormData({ 
        title: "", 
        description: "", 
        goalType: "CUSTOM", 
        targetLevel: "", 
        targetCount: "", 
        targetDate: "" 
      })
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsSubmitting(true)
    try {
      const payload = {
        title: formData.title,
        description: formData.description || null,
        goalType: formData.goalType,
        targetLevel: formData.targetLevel || null,
        targetCount: formData.targetCount ? parseInt(formData.targetCount) : null,
        // Convert YYYY-MM-DD to ISO string if needed, or backend might accept YYYY-MM-DD directly
        targetDate: formData.targetDate ? new Date(formData.targetDate).toISOString() : null
      }
      
      if (editId) {
        await api.put(`/goals/${editId}`, payload)
      } else {
        await api.post("/goals", payload)
      }
      setIsDialogOpen(false)
      setEditId(null)
      fetchGoals()
    } catch (err) {
      console.error("Failed to save", err)
    } finally {
      setIsSubmitting(false)
    }
  }

  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'COMPLETED':
        return <CheckCircle2 className="h-5 w-5 text-green-500" />
      case 'FAILED':
        return <XCircle className="h-5 w-5 text-red-500" />
      default:
        return <Clock className="h-5 w-5 text-amber-500" />
    }
  }

  const getStatusText = (status: string) => {
    switch (status) {
      case 'COMPLETED': return 'Completed'
      case 'FAILED': return 'Failed'
      default: return 'In Progress'
    }
  }

  return (
    <div className="container max-w-6xl mx-auto py-8 px-4 space-y-8">
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 className="text-3xl font-bold tracking-tight flex items-center gap-2">
            <Target className="h-8 w-8 text-primary" />
            Learning Goals
          </h1>
          <p className="text-muted-foreground mt-1">Set and track your Japanese learning milestones.</p>
        </div>
        
        <Dialog open={isDialogOpen} onOpenChange={handleOpenChange}>
          <DialogTrigger asChild>
            <Button size="lg" className="shadow-sm">
              <Plus className="mr-2 h-5 w-5" /> Add New Goal
            </Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[500px]">
            <DialogHeader>
              <DialogTitle>{editId ? "Edit Goal" : "Create New Goal"}</DialogTitle>
              <DialogDescription>
                Define what you want to achieve and set a deadline.
              </DialogDescription>
            </DialogHeader>
            <form onSubmit={handleSubmit} className="space-y-4 py-4">
              <div className="space-y-2">
                <Label htmlFor="title">Goal Title <span className="text-destructive">*</span></Label>
                <Input 
                  id="title" 
                  value={formData.title} 
                  onChange={(e) => setFormData({...formData, title: e.target.value})} 
                  placeholder="e.g. Pass JLPT N4, Learn 500 Kanji"
                  required 
                />
              </div>
              
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="goalType">Goal Type</Label>
                  <select 
                    id="goalType"
                    className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                    value={formData.goalType}
                    onChange={(e) => setFormData({...formData, goalType: e.target.value})}
                  >
                    <option value="CUSTOM">Custom</option>
                    <option value="JLPT">JLPT Test</option>
                    <option value="VOCABULARY">Vocabulary</option>
                    <option value="KANJI">Kanji</option>
                    <option value="GRAMMAR">Grammar</option>
                  </select>
                </div>
                
                <div className="space-y-2">
                  <Label htmlFor="targetDate">Target Date <span className="text-destructive">*</span></Label>
                  <Input 
                    id="targetDate" 
                    type="date"
                    value={formData.targetDate} 
                    onChange={(e) => setFormData({...formData, targetDate: e.target.value})} 
                    required 
                  />
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="targetLevel">Target Level</Label>
                  <Input 
                    id="targetLevel" 
                    placeholder="e.g. N5, N4"
                    value={formData.targetLevel} 
                    onChange={(e) => setFormData({...formData, targetLevel: e.target.value})} 
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="targetCount">Target Count</Label>
                  <Input 
                    id="targetCount" 
                    type="number"
                    min="1"
                    placeholder="e.g. 100"
                    value={formData.targetCount} 
                    onChange={(e) => setFormData({...formData, targetCount: e.target.value})} 
                  />
                </div>
              </div>

              <div className="space-y-2">
                <Label htmlFor="description">Description (Optional)</Label>
                <Textarea 
                  id="description" 
                  value={formData.description} 
                  onChange={(e) => setFormData({...formData, description: e.target.value})} 
                  placeholder="Details about your goal..."
                  rows={3}
                />
              </div>
              
              <DialogFooter className="pt-4">
                <Button type="submit" disabled={isSubmitting} className="w-full sm:w-auto">
                  {isSubmitting && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                  {editId ? "Update Goal" : "Save Goal"}
                </Button>
              </DialogFooter>
            </form>
          </DialogContent>
        </Dialog>
      </div>

      {loading ? (
        <div className="flex justify-center py-20">
          <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
        </div>
      ) : goals.length === 0 ? (
        <Card className="border-dashed bg-transparent shadow-none">
          <CardContent className="flex flex-col items-center justify-center py-20 text-center">
            <Target className="h-12 w-12 text-muted-foreground/50 mb-4" />
            <h3 className="text-xl font-medium mb-2">No goals set yet</h3>
            <p className="text-muted-foreground max-w-md mb-6">
              Setting clear goals is the first step to mastering Japanese. Create your first goal to track your progress!
            </p>
            <Button onClick={() => setIsDialogOpen(true)}>Create a Goal</Button>
          </CardContent>
        </Card>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {goals.map((goal) => (
            <Card key={goal.id} className="flex flex-col overflow-hidden transition-all hover:shadow-md border-border/50">
              <CardHeader className="pb-4">
                <div className="flex justify-between items-start gap-4">
                  <div>
                    <div className="flex items-center gap-2 mb-1">
                      <span className="text-xs font-semibold px-2 py-0.5 rounded-full bg-primary/10 text-primary">
                        {goal.goalType}
                      </span>
                      {goal.targetLevel && (
                        <span className="text-xs font-semibold px-2 py-0.5 rounded-full bg-secondary text-secondary-foreground">
                          {goal.targetLevel}
                        </span>
                      )}
                    </div>
                    <CardTitle className="line-clamp-2 leading-tight text-lg">{goal.title}</CardTitle>
                  </div>
                  <div className="flex -mr-2">
                    <Button variant="ghost" size="icon" className="h-8 w-8 text-muted-foreground" onClick={() => openEdit(goal)}>
                      <Edit className="h-4 w-4" />
                    </Button>
                    <Button variant="ghost" size="icon" className="h-8 w-8 text-muted-foreground hover:text-destructive" onClick={() => deleteGoal(goal.id)}>
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </div>
                </div>
                {goal.description && (
                  <CardDescription className="line-clamp-2 mt-2">{goal.description}</CardDescription>
                )}
              </CardHeader>
              
              <CardContent className="flex-1 pb-4">
                {goal.targetCount !== null && goal.targetCount !== undefined && goal.targetCount > 0 && (
                  <div className="space-y-2 mt-2">
                    <div className="flex justify-between text-sm">
                      <span className="text-muted-foreground">Progress</span>
                      <span className="font-medium">{goal.currentCount || 0} / {goal.targetCount}</span>
                    </div>
                    <div className="h-2 w-full bg-secondary rounded-full overflow-hidden">
                      <div 
                        className="h-full bg-primary transition-all duration-500 ease-in-out" 
                        style={{ width: `${Math.min(100, Math.max(0, goal.progressPercentage || 0))}%` }}
                      />
                    </div>
                  </div>
                )}
              </CardContent>
              
              <CardFooter className="bg-muted/30 border-t py-3 flex justify-between items-center text-sm">
                <div className="flex items-center gap-1.5 text-muted-foreground" title="Target Date">
                  <Calendar className="h-4 w-4" />
                  <span>{new Date(goal.targetDate).toLocaleDateString(undefined, { month: 'short', day: 'numeric', year: 'numeric' })}</span>
                </div>
                <div className="flex items-center gap-1.5 font-medium" title={`Status: ${getStatusText(goal.status)}`}>
                  {getStatusIcon(goal.status)}
                  <span className="hidden sm:inline-block">{getStatusText(goal.status)}</span>
                </div>
              </CardFooter>
            </Card>
          ))}
        </div>
      )}
    </div>
  )
}
