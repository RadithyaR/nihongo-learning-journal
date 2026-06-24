"use client"

import { useEffect, useState } from "react"
import { api } from "@/lib/axios"
import { Plus, Search, Edit, Trash2, Heart, Loader2 } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
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

interface Kanji {
  id: string
  character: string
  meaning: string
  onyomi?: string
  kunyomi?: string
  jlpt_level?: string
  status?: string
  favourite: boolean
}

export default function KanjiPage() {
  const [kanjis, setKanjis] = useState<Kanji[]>([])
  const [loading, setLoading] = useState(true)
  const [search, setSearch] = useState("")

  const [isDialogOpen, setIsDialogOpen] = useState(false)
  const [editId, setEditId] = useState<string | null>(null)
  const [formData, setFormData] = useState({
    character: "",
    meaning: "",
    onyomi: "",
    kunyomi: "",
    jlpt_level: ""
  })
  const [isSubmitting, setIsSubmitting] = useState(false)

  const fetchKanjis = async (query = "") => {
    setLoading(true)
    try {
      const response = await api.get(`/kanjis${query ? `?search=${query}` : ""}`)
      setKanjis(response.data.data)
    } catch (err) {
      console.error("Failed to load kanjis", err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    const delayDebounceFn = setTimeout(() => {
      fetchKanjis(search)
    }, 500)
    return () => clearTimeout(delayDebounceFn)
  }, [search])

  const toggleFavourite = async (id: string) => {
    try {
      await api.patch(`/kanjis/${id}/favourite`)
      setKanjis(kanjis.map(k => 
        k.id === id ? { ...k, favourite: !k.favourite } : k
      ))
    } catch (err) {
      console.error("Failed to toggle favourite", err)
    }
  }

  const deleteKanji = async (id: string) => {
    if (!confirm("Are you sure you want to delete this item?")) return
    
    try {
      await api.delete(`/kanjis/${id}`)
      setKanjis(kanjis.filter(k => k.id !== id))
    } catch (err) {
      console.error("Failed to delete", err)
    }
  }

  const openEdit = (kanji: Kanji) => {
    setFormData({
      character: kanji.character,
      meaning: kanji.meaning,
      onyomi: kanji.onyomi || "",
      kunyomi: kanji.kunyomi || "",
      jlpt_level: kanji.jlpt_level || ""
    })
    setEditId(kanji.id)
    setIsDialogOpen(true)
  }

  const handleOpenChange = (open: boolean) => {
    setIsDialogOpen(open)
    if (!open) {
      setEditId(null)
      setFormData({ character: "", meaning: "", onyomi: "", kunyomi: "", jlpt_level: "" })
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsSubmitting(true)
    try {
      if (editId) {
        await api.put(`/kanjis/${editId}`, formData)
      } else {
        await api.post("/kanjis", formData)
      }
      setIsDialogOpen(false)
      setEditId(null)
      setFormData({ character: "", meaning: "", onyomi: "", kunyomi: "", jlpt_level: "" })
      fetchKanjis(search)
    } catch (err) {
      console.error("Failed to save", err)
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <div className="container max-w-6xl mx-auto py-8 px-4 space-y-6">
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Kanji</h1>
          <p className="text-muted-foreground">Manage your Japanese Kanji collection.</p>
        </div>
        
        <Dialog open={isDialogOpen} onOpenChange={handleOpenChange}>
          <DialogTrigger asChild>
            <Button>
              <Plus className="mr-2 h-4 w-4" /> Add New Kanji
            </Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[500px]">
            <DialogHeader>
              <DialogTitle>{editId ? "Edit Kanji" : "Add Kanji"}</DialogTitle>
              <DialogDescription>
                {editId ? "Edit your Kanji details." : "Add a new Japanese Kanji to your study collection."}
              </DialogDescription>
            </DialogHeader>
            <form onSubmit={handleSubmit} className="space-y-4 py-4">
              <div className="space-y-2">
                <Label htmlFor="character">Kanji Character <span className="text-destructive">*</span></Label>
                <Input 
                  id="character" 
                  value={formData.character} 
                  onChange={(e) => setFormData({...formData, character: e.target.value})} 
                  required 
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="meaning">Meaning <span className="text-destructive">*</span></Label>
                <Input 
                  id="meaning" 
                  value={formData.meaning} 
                  onChange={(e) => setFormData({...formData, meaning: e.target.value})} 
                  required
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="onyomi">Onyomi (Chinese Reading)</Label>
                <Input 
                  id="onyomi" 
                  value={formData.onyomi} 
                  onChange={(e) => setFormData({...formData, onyomi: e.target.value})} 
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="kunyomi">Kunyomi (Japanese Reading)</Label>
                <Input 
                  id="kunyomi" 
                  value={formData.kunyomi} 
                  onChange={(e) => setFormData({...formData, kunyomi: e.target.value})} 
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="jlpt_level">JLPT Level</Label>
                <Input 
                  id="jlpt_level" 
                  value={formData.jlpt_level} 
                  onChange={(e) => setFormData({...formData, jlpt_level: e.target.value})} 
                  placeholder="e.g. N5, N4"
                />
              </div>
              <DialogFooter>
                <Button type="submit" disabled={isSubmitting}>
                  {isSubmitting && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                  Save Item
                </Button>
              </DialogFooter>
            </form>
          </DialogContent>
        </Dialog>
      </div>

      <div className="flex items-center relative max-w-sm">
        <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
        <Input 
          placeholder="Search kanji..." 
          className="pl-9"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
        />
      </div>

      <div className="border rounded-md">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[50px]"></TableHead>
              <TableHead>Character</TableHead>
              <TableHead>Meaning</TableHead>
              <TableHead>Onyomi</TableHead>
              <TableHead>Kunyomi</TableHead>
              <TableHead>Status</TableHead>
              <TableHead className="text-right">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {loading ? (
              <TableRow>
                <TableCell colSpan={7} className="text-center py-10 text-muted-foreground">
                  <Loader2 className="h-6 w-6 animate-spin mx-auto" />
                </TableCell>
              </TableRow>
            ) : kanjis.length === 0 ? (
              <TableRow>
                <TableCell colSpan={7} className="text-center py-10 text-muted-foreground">
                  No Kanji found. Add your first Kanji!
                </TableCell>
              </TableRow>
            ) : (
              kanjis.map((item) => (
                <TableRow key={item.id}>
                  <TableCell>
                    <button 
                      onClick={() => toggleFavourite(item.id)}
                      className={`hover:bg-muted p-1.5 rounded-full transition-colors ${item.favourite ? 'text-red-500' : 'text-muted-foreground'}`}
                    >
                      <Heart className="h-4 w-4" fill={item.favourite ? "currentColor" : "none"} />
                    </button>
                  </TableCell>
                  <TableCell className="font-bold text-2xl">{item.character}</TableCell>
                  <TableCell>{item.meaning}</TableCell>
                  <TableCell className="text-muted-foreground">{item.onyomi || "-"}</TableCell>
                  <TableCell className="text-muted-foreground">{item.kunyomi || "-"}</TableCell>
                  <TableCell>
                    <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-primary/10 text-primary">
                      {item.status || "NEW"}
                    </span>
                  </TableCell>
                  <TableCell className="text-right">
                    <Button variant="ghost" size="icon" title="Edit" onClick={() => openEdit(item)}>
                      <Edit className="h-4 w-4 text-muted-foreground" />
                    </Button>
                    <Button variant="ghost" size="icon" title="Delete" onClick={() => deleteKanji(item.id)}>
                      <Trash2 className="h-4 w-4 text-destructive" />
                    </Button>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </div>
    </div>
  )
}
