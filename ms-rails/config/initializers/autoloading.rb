module Services; end

Rails.autoloaders.main.push_dir("#{Rails.root}/app/services",namespace: Services)