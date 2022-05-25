from distutils.text_file import TextFile
from django.shortcuts import render, redirect
from .models import Text
from .forms import TextForm


def index(request):
    error = ''
    if request.method == 'POST':
        form = TextForm(request.POST)
        if form.is_valid():
            form.save()
            return redirect('home')
        else:
            error = 'Форма была неверной'


    form = TextForm()
    texts = Text.objects.all()
    return render(request, 'main/index.html', {'title': 'Главная страни сайта', 'texts': texts, 'form': form, 'error': error})


def about(request):
    return render(request, 'main/about.html')