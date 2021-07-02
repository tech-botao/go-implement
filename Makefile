test:
	echo "test something"
cobra-gen:
	cobra init --pkg-name github.com/tech-botao/go-implement/app/cobra-gen app/cobra-gen
cd/cobra-gen:
	cd app/cobra-gen
cobra-add:cd/cobra-gen
	cobra add config -a zhoubo
